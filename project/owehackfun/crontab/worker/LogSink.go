package worker

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golearn/project/owehackfun/crontab/common"
	"time"
)

// mongodb 存储日志

type LogSink struct {
	client         *mongo.Client
	logCollection  *mongo.Collection
	logChan        chan *common.JobLog
	autoCommitChan chan *common.LogBatch
}

var (
	// 单例
	G_logSink *LogSink
)

func (logSink *LogSink) saveLogs(batch *common.LogBatch) {
	logSink.logCollection.InsertMany(context.TODO(), batch.Logs)
}
func (logSink *LogSink) writeLoop() {
	var (
		log          *common.JobLog
		logBatch     *common.LogBatch
		commitTimer  *time.Timer
		timeoutBatch *common.LogBatch // 超时批次
	)

	for {
		select {
		case log = <-logSink.logChan:
			// 把这条log写到 mongodb中
			// logSink.logCollection.InsertOne
			// 每次插入需要等待 mongodb的一次请求往返，耗时可能因为网络慢话费较长时间
			if logBatch == nil {
				logBatch = &common.LogBatch{}
				// 让这个批次超时自动提交（给1秒的时间）
				commitTimer = time.AfterFunc(
					time.Duration(G_config.JobLogCommitTimeout),
					func(batch *common.LogBatch) func() {
						return func() {
							logSink.autoCommitChan <- batch
						}
					}(logBatch),
				)
			}

			// 把新的日志追加到批次中
			logBatch.Logs = append(logBatch.Logs, log)

			// 如果批次满了，就立即发送
			if len(logBatch.Logs) >= G_config.JobLogBatchSize {
				// 发送日志
				logSink.saveLogs(logBatch)

				// 清空 logBatch
				logBatch = nil

				// 取消定时器
				commitTimer.Stop()
			}
		case timeoutBatch = <-logSink.autoCommitChan: // 过期的批次
			// 判断过期批次是否仍旧是当前的批次
			if timeoutBatch != logBatch {
				continue // 跳过已经被提交的批次
			}

			// 把批次写入到mongo 中
			logSink.saveLogs(timeoutBatch)
			logBatch = nil
		}

	}
}

func InitLogSink() (err error) {
	var (
		client *mongo.Client
		ctx    context.Context
	)

	// 建立 mongodb 连接
	ctx, _ = context.WithTimeout(context.Background(), time.Duration(G_config.MongodbConnectTimeOut)*time.Second)
	if client, err = mongo.Connect(ctx, options.Client().ApplyURI(G_config.MongodbUri)); err != nil {
		return
	}

	// 选择 db 和 collection
	G_logSink = &LogSink{
		client:         client,
		logCollection:  client.Database("cron").Collection("log"),
		logChan:        make(chan *common.JobLog, 1000),
		autoCommitChan: make(chan *common.LogBatch, 1000),
	}

	// 启动一个mongodb 处理协程
	go G_logSink.writeLoop()
	return
}

// 发送日志
func (logSink *LogSink) Append(jobLog *common.JobLog) {
	select {
	case logSink.logChan <- jobLog:
	default:
		// 队列满了就丢弃
	}
}

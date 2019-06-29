package master

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golearn/project/owehackfun/crontab/common"
	"time"
)

// mongodb 日志管理
type LogMgr struct {
	client        *mongo.Client
	logCollection *mongo.Collection
}

var (
	G_logMgr *LogMgr
)

func InitLogMgr() (err error) {
	var (
		client *mongo.Client
		ctx    context.Context
	)

	// 建立 mongodb 连接
	ctx, _ = context.WithTimeout(context.Background(), time.Duration(G_config.MongodbConnectTimeOut)*time.Second)
	if client, err = mongo.Connect(ctx, options.Client().ApplyURI(G_config.MongodbUri)); err != nil {
		return
	}

	G_logMgr = &LogMgr{
		client:        client,
		logCollection: client.Database("cron").Collection("log"),
	}
	return
}

func (LogMgr *LogMgr) ListLog(name string, skip, limit int64) (logArr []*common.JobLog, err error) {
	var (
		filter  *common.JobLogFilter
		logSort *common.SortLogByStartTime
		findopt *options.FindOptions
		cursor  *mongo.Cursor
		jogLog  *common.JobLog
	)

	// len(logArr)
	logArr = make([]*common.JobLog, 0)

	filter = &common.JobLogFilter{JobName: name}

	// 按照任务开始时间倒排
	logSort = &common.SortLogByStartTime{SortOrder: -1}

	findopt = &options.FindOptions{
		Sort:  logSort,
		Skip:  &skip,
		Limit: &limit,
	}

	// 查询
	if cursor, err = LogMgr.logCollection.Find(context.TODO(), filter, findopt); err != nil {
		return
	}
	// 延迟释放游标
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		jogLog = &common.JobLog{}

		// 反序列化BSON
		if err = cursor.Decode(jogLog); err != nil {
			continue // 有日志不合法
		}

		logArr = append(logArr, jogLog)
	}

	return
}

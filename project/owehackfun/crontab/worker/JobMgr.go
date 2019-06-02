package worker

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
	"golearn/project/owehackfun/crontab/common"
	"time"
)

// 任务管理器
type JobMgr struct {
	client  *clientv3.Client
	kv      clientv3.KV
	lease   clientv3.Lease
	watcher clientv3.Watcher
}

var (
	// 单例
	G_jobMgr *JobMgr
)

// 监听任务变化
func (jobMgr *JobMgr) watchJobs() (err error) {
	var (
		getResp            *clientv3.GetResponse
		kvpair             *mvccpb.KeyValue
		job                *common.Job
		watchStartRevision int64
		watchChan          clientv3.WatchChan
		watchResp          clientv3.WatchResponse
		watchEvent         *clientv3.Event
		jobName            string
		jobEvent           *common.JobEvent
	)

	// 1. get 一下 /cron/jobs/目录下的所有任务，并且获知当前集群的 revision
	if getResp, err = jobMgr.kv.Get(context.TODO(), common.JOB_SAVE_DIR, clientv3.WithPrefix()); err != nil {
		return
	}

	// 当前有哪些任务
	for _, kvpair = range getResp.Kvs {
		// 反序列化 json 得到 Job
		if job, err = common.UnpackJob(kvpair.Value); err == nil {
			jobEvent = common.BuildJobEvent(common.JOB_EVENT_SAVE, job)
			// TODO：把这个 job 同步给 scheduler（调度协程）
			fmt.Println(*jobEvent)
		}
	}

	// 2. 从该 revision 向后监听变化事件
	go func() { // 监听协程
		// 从GET 时刻的后续版本开始监听变化
		watchStartRevision = getResp.Header.Revision + 1
		// 监听 /cron/jobs? 目录后续变化
		watchChan = jobMgr.watcher.Watch(context.TODO(), common.JOB_SAVE_DIR, clientv3.WithRev(watchStartRevision), clientv3.WithPrefix())
		// 处理监听事件
		for watchResp = range watchChan {
			for _, watchEvent = range watchResp.Events {
				switch watchEvent.Type {
				case mvccpb.PUT: // 保存任务事件
					if job, err = common.UnpackJob(watchEvent.Kv.Value); err != nil {
						continue
					}
					// 构造一个Event 事件
					jobEvent = common.BuildJobEvent(common.JOB_EVENT_SAVE, job)

					// TODO: 反序列化 Job， 推送更新事件给scheduler
				case mvccpb.DELETE: // 任务被删除
					// Delete /cron/jobs/job10
					// TODO：推送删除事件给 scheduler
					jobName = common.ExtractJobName(string(watchEvent.Kv.Key))

					job = &common.Job{Name: jobName}

					// 构造一个删除 Event
					jobEvent = common.BuildJobEvent(common.JOB_EVENT_DELETE, job)

				}
				// TODO：推送删除事件给 scheduler
				// G_Scheduler.PushJobEvent(jobEvent)
				fmt.Println(*jobEvent)

			}

		}
	}()

	return

}

// 初始化管理
func InitJobMgr() (err error) {
	var (
		config  clientv3.Config
		client  *clientv3.Client
		kv      clientv3.KV
		lease   clientv3.Lease
		watcher clientv3.Watcher
	)

	// 初始化配置
	config = clientv3.Config{
		Endpoints:   G_config.EtcdEndpoints,                                    // 集群地址
		DialTimeout: time.Duration(G_config.EtcdDialTmeOut) * time.Millisecond, // 连接超时
	}

	// 建立连接
	if client, err = clientv3.New(config); err != nil {
		return
	}

	// 得到 KV 和 Lease 的 API 子集
	kv = clientv3.NewKV(client)
	lease = clientv3.NewLease(client)
	watcher = clientv3.NewWatcher(client)

	// 赋值单例
	G_jobMgr = &JobMgr{
		client:  client,
		kv:      kv,
		lease:   lease,
		watcher: watcher,
	}

	// 启动监听
	G_jobMgr.watchJobs()

	return

}

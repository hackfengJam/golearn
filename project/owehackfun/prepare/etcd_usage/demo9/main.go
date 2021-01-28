package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	// 乐观锁

	var (
		config         clientv3.Config
		client         *clientv3.Client
		err            error
		key            string
		lease          clientv3.Lease
		leaseGrantResp *clientv3.LeaseGrantResponse
		leaseId        clientv3.LeaseID
		keepRespChan   <-chan *clientv3.LeaseKeepAliveResponse
		keepResp       *clientv3.LeaseKeepAliveResponse
		ctx            context.Context
		cancelFunc     context.CancelFunc
		kv             clientv3.KV
		txn            clientv3.Txn
		txnResp        *clientv3.TxnResponse
	)

	// 客户端配置
	config = clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"}, // 集群列表
		DialTimeout: 5 * time.Second,
	}

	// 建立连接
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	// lease 实现锁自动过期：
	// op 操作
	// txn 事务：if else then

	// 1. 上锁（创建租约，自动续约，拿着租约去抢占一个key）

	// 申请一个 lease（租约）
	lease = clientv3.NewLease(client)

	// 申请一个 5 秒的租约
	if leaseGrantResp, err = lease.Grant(context.TODO(), 5); err != nil {
		fmt.Println(err)
		return
	}

	// 拿到租约ID
	leaseId = leaseGrantResp.ID

	// 准备一个用于取消自动续租的context
	ctx, cancelFunc = context.WithCancel(context.TODO())

	// 确保函数退出后，自动续约会停止
	defer cancelFunc()
	defer lease.Revoke(context.TODO(), leaseId)

	// 自动续租
	if keepRespChan, err = lease.KeepAlive(ctx, leaseId); err != nil {
		fmt.Println(err)
		return
	}

	// 处理续约应答的协程
	go func() {
		for {
			select {
			case keepResp = <-keepRespChan:
				if keepResp == nil {
					fmt.Println("租约已失效")
					goto END
				} else { // 每秒会续约一次，所以会收到一次应答
					fmt.Println("收到自动续约应答", keepResp.ID)
				}
			}
		}
	END:
	}()

	// 抢Key：if 不存在key,then 设置它，else 抢锁失败
	kv = clientv3.NewKV(client)

	// 创建事务
	txn = kv.Txn(context.TODO())

	// 定义事务

	// 如果key不存在
	key = "/cron/lock/job9"
	txn.If(clientv3.Compare(clientv3.CreateRevision(key), "=", 0)).
		Then(clientv3.OpPut(key, "xxx", clientv3.WithLease(leaseId))).
		Else(clientv3.OpGet(key)) // 否则抢锁失败

	// 提交事务
	if txnResp, err = txn.Commit(); err != nil {
		fmt.Println(err)
		return
	}

	// 判断是否抢到了锁
	if !txnResp.Succeeded {
		fmt.Println("锁被占用", string(txnResp.Responses[0].GetResponseRange().Kvs[0].Value))
		return
	}

	// 2. 处理业务
	fmt.Println("处理任务")
	time.Sleep(5 * time.Second)

	// 在锁内，很安全

	// 3. 释放锁（取消自动续约，释放租约）

	// defer 会租约释放，关联的KV就被删除了
}

package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) OutPut(items []interface{}) {
	// TODO to mq
	itemCount := 0

	for _, item := range items {
		log.Printf("Got item: #%d %v", itemCount, item)
		itemCount++
	}

}

func (e *ConcurrentEngine) Run(seeds ...Request) {
    // worker -> 引擎 的 chan
	out := make(chan ParseResult)

	// 启动 Scheduler
	e.Scheduler.Run()

	// 启动 WorkerCount 个 worker
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	// 将所有种子 Submit
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	for {
		// 获取结果
		result := <-out

		// 输出
		e.OutPut(result.Items)

		// 将结果中的 Requests Submit
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}


func createWorker(in chan Request,
	out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			// tell scheduler i'm ready
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

package engine

import (
	"log"
)

type ConcurrentRequestEngine struct {
	Scheduler   ConcurrentRequestScheduler
	WorkerCount int
}


type ConcurrentRequestScheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
	//WorkerReady(chan Request)
	//Run()
}

func (e ConcurrentRequestEngine) OutPut(items []interface{}) {
	// TODO to mq
	itemCount := 0

	for _, item := range items {
		log.Printf("Got item: #%d %v", itemCount, item)
		itemCount++
	}


}

func (e *ConcurrentRequestEngine) Run(seeds ...Request) {

	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigureMasterWorkerChan(in)

	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(in, out)
	}

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

func (e *ConcurrentRequestEngine) createWorker(
	in chan Request, out chan ParseResult) {
	go func() {
		for {
			// tell scheduler i'm ready
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

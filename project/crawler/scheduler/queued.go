package scheduler

import (
	"golearn/project/crawler/engine"
)

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	// 每个 worker 都有自己的 chan
	return make(chan engine.Request)
}

func (s *QueuedScheduler) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}
			select {

			case r := <-s.requestChan:
				// send r to a ?worker
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				// send r next_request to w
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				// 当真正调度成功，则出队
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()

}

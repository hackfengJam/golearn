package scheduler

import "learngo/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(
	c chan engine.Request) {
	s.workerChan = c
}

func (s *SimpleScheduler) Submit(
	r engine.Request) {
	// send request down to worker chan
	go func() { s.workerChan <- r }()
}

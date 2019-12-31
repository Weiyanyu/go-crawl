package scheduler

import "go-crawl/engine"

type SimpleScheduler struct {
	workerChannel chan engine.Request
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	go func(r engine.Request) {
		s.workerChannel <- request
	}(request)

}

func (s *SimpleScheduler) WorkerReady(worker chan engine.Request) {
}

func (s *SimpleScheduler) WorkerChannel() chan engine.Request {
	return s.workerChannel
}

func (s *SimpleScheduler) Run() {
	s.workerChannel = make(chan engine.Request)
}

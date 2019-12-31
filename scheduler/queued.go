package scheduler

import "go-crawl/engine"

type QueuedScheduler struct {
	requestChannel chan engine.Request
	workerChannel  chan chan engine.Request
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChannel <- r
}

func (s *QueuedScheduler) WorkerReady(worker chan engine.Request) {
	s.workerChannel <- worker
}

func (s *QueuedScheduler) WorkerChannel() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueuedScheduler) Run() {
	s.requestChannel = make(chan engine.Request)
	s.workerChannel = make(chan chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case r := <-s.requestChannel:
				requestQ = append(requestQ, r)
			case w := <-s.workerChannel:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}

	}()

}

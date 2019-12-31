package engine

import "log"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

func (e *ConcurrentEngine) Run(Seeds ...Request) {

	out := make(chan ParseResult)
	e.Scheduler.Run()
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChannel(), out, e.Scheduler)
	}
	count := 0
	for _, r := range Seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out

		for _, item := range result.Items {
			count++
			log.Printf("Got item : #%d : %v\n", count, item)
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}

	}
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
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

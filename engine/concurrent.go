package engine

type ConcurrentEngine struct {
	Scheduler        Scheduler
	WorkerCount      int
	ItemChanel       chan Item
	RequestProcessor RequestProcessor
}

type RequestProcessor func(Request) (ParseResult, error)

func (e *ConcurrentEngine) Run(Seeds ...Request) {

	out := make(chan ParseResult)
	e.Scheduler.Run()
	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.WorkerChannel(), out, e.Scheduler)
	}
	for _, r := range Seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out

		for _, item := range result.Items {
			go func(i Item) {
				e.ItemChanel <- i
			}(item)
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}

	}
}

func (e *ConcurrentEngine) createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := e.RequestProcessor(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

package main

import (
	"go-crawl/engine"
	"go-crawl/scheduler"
	"go-crawl/zhenai/parser"
)

func main() {
	e := &engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 20,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

}

package main

import (
	"go-crawl/engine"
	"go-crawl/persist"
	"go-crawl/scheduler"
	"go-crawl/zhenai/parser"
)

func main() {
	itemChannel, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}
	e := &engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 20,
		ItemChanel:  itemChannel,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

}

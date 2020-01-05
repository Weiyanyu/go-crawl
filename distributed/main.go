package main

import (
	"fmt"
	"go-crawl/distributed/config"
	"go-crawl/distributed/persist/client"
	"go-crawl/distributed/persist/server"

	workerclient "go-crawl/distributed/worker/client"
	workerserver "go-crawl/distributed/worker/server"
	"go-crawl/engine"
	"go-crawl/scheduler"
	"go-crawl/zhenai/parser"
)

func main() {

	serverNotifier := make(chan struct{})
	go server.ServrRpc(fmt.Sprintf(":%d", config.ItemSaverPort), config.ItemSaverESIndex, serverNotifier)

	<-serverNotifier
	itemChannel, err := client.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort), "dating_profile")
	if err != nil {
		panic(err)
	}
	go workerserver.ServeRpc(fmt.Sprintf(":%d", config.WorkerPort0), serverNotifier)

	<-serverNotifier

	requestProcessor, err := workerclient.CreateProcessor()
	if err != nil {
		panic(err)
	}

	e := &engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      20,
		ItemChanel:       itemChannel,
		RequestProcessor: requestProcessor,
	}
	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})

}

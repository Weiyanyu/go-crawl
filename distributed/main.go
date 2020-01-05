package main

import (
	"fmt"
	"go-crawl/distributed/config"
	"go-crawl/distributed/persist/client"
	"go-crawl/distributed/persist/server"
	"go-crawl/engine"
	"go-crawl/scheduler"
	"go-crawl/zhenai/parser"
	"time"
)

func main() {
	serverNotifier := make(chan struct{})
	go server.ServrRpc(fmt.Sprintf(":%d", config.ItemSaverPort), config.ItemSaverESIndex, serverNotifier)

	<-serverNotifier

	time.Sleep(1 * time.Second)
	itemChannel, err := client.ItemSaver(":1234", "dating_profile")
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

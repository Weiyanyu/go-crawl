package main

import (
	"fmt"
	"go-crawl/distributed/config"
	"go-crawl/distributed/persist/client"
	"go-crawl/distributed/rpcsupport"
	"log"
	"net/rpc"

	workerclient "go-crawl/distributed/worker/client"
	"go-crawl/engine"
	"go-crawl/scheduler"
	"go-crawl/zhenai/parser"
)

func main() {

	itemChannel, err := client.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverPort), "dating_profile")
	if err != nil {
		panic(err)
	}

	pool := createProcessorPool(config.WorkerHostList)
	requestProcessor := workerclient.CreateProcessor(pool)

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

func createProcessorPool(hosts []string) chan *rpc.Client {

	var clients []*rpc.Client
	for _, host := range hosts {
		client, err := rpcsupport.NewRpcClient(host)
		if err != nil {
			log.Printf("Error connection to %s : %v", host, err)
		} else {
			clients = append(clients, client)
			log.Printf("Connected to %s", host)
		}
	}

	pool := make(chan *rpc.Client)

	go func() {
		for {
			for _, client := range clients {
				pool <- client
			}
		}
	}()

	return pool
}

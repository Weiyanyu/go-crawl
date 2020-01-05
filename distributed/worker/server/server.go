package server

import (
	"go-crawl/distributed/rpcsupport"
	"go-crawl/distributed/worker"
	"log"
)

func ServeRpc(host string, serverNotifier chan struct{}) {
	log.Fatal(rpcsupport.ServeRpc(host, worker.CrawlService{}, serverNotifier))
}

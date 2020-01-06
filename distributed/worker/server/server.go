package main

import (
	"flag"
	"fmt"
	"go-crawl/distributed/rpcsupport"
	"go-crawl/distributed/worker"
	"log"
)

var port = flag.Int("port", 0, "worker server port")

func main() {
	flag.Parse()
	if *port == 0 {
		panic("port can't equals zero")
	}
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", *port), worker.CrawlService{}))
}

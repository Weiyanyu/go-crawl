package server

import (
	"go-crawl/distributed/config"
	"go-crawl/distributed/rpcsupport"
	"go-crawl/distributed/worker"
	"testing"
)

func TestWorkerServer(t *testing.T) {
	const host = ":9000"
	ch := make(chan struct{})

	go ServeRpc(host, ch)

	<-ch

	client, err := rpcsupport.NewRpcClient(host)
	if err != nil {
		panic(err)
	}

	request := worker.Request{
		Url: "http://www.zhenai.com/zhenghun/bishan",
		Parser: worker.SerializedParser{
			Name: config.ParseCityAndProfile,
		},
	}
	var result worker.ParseResult
	client.Call(config.WorkerRpcFuncName, request, &result)
	fmt.Println(result)
}

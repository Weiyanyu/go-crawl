package client

import (
	"fmt"
	"go-crawl/distributed/config"
	"go-crawl/distributed/rpcsupport"
	"go-crawl/distributed/worker"
	"go-crawl/engine"
)

func CreateProcessor() (engine.RequestProcessor, error) {
	client, err := rpcsupport.NewRpcClient(fmt.Sprintf(":%d", config.WorkerPort0))
	if err != nil {
		return nil, err
	}
	return func(req engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializeRequest(req)
		var sResult worker.ParseResult
		err := client.Call(config.WorkerRpcFuncName, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult), nil
	}, nil
}

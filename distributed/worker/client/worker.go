package client

import (
	"go-crawl/distributed/config"
	"go-crawl/distributed/worker"
	"go-crawl/engine"
	"net/rpc"
)

func CreateProcessor(clientCh chan *rpc.Client) engine.RequestProcessor {
	return func(req engine.Request) (engine.ParseResult, error) {

		sReq := worker.SerializeRequest(req)
		var sResult worker.ParseResult
		client := <-clientCh
		err := client.Call(config.WorkerRpcFuncName, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult), nil
	}
}

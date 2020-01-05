package client

import (
	"go-crawl/distributed/config"
	"go-crawl/distributed/rpcsupport"
	"go-crawl/engine"
	"log"
)

func ItemSaver(host, indexName string) (chan engine.Item, error) {
	client, err := rpcsupport.NewRpcClient(host)

	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		count := 0
		for {
			item := <-out
			count++
			var result string
			err = client.Call(config.ItemSaverRpcFuncName, item, &result)
			if err != nil {
				log.Printf("Item Saver : occure error : %v", err)
				continue
			}
		}
	}()
	return out, nil
}

package server

import (
	"go-crawl/distributed/persist"
	"go-crawl/distributed/rpcsupport"
	"log"

	"github.com/olivere/elastic"
)

func ServrRpc(host string, index string, serverNotifier chan struct{}) {
	esClient, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	log.Fatal(rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: esClient,
		Index:  index,
	}, serverNotifier))

}

package server

import (
	"go-crawl/distributed/persist"
	"go-crawl/distributed/rpcsupport"

	"github.com/olivere/elastic"
)

func ServrRpc(host string, index string, serverNotifier chan struct{}) error {
	esClient, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: esClient,
		Index:  index,
	}, serverNotifier)
}

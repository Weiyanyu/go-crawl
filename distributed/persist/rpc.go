package persist

import (
	"go-crawl/engine"
	"go-crawl/persist"
	"log"

	"github.com/olivere/elastic"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (service *ItemSaverService) Save(item engine.Item, result *string) error {
	log.Printf("ItemSaverService : Save Item : %v", item)

	err := persist.Save(service.Client, service.Index, item)
	if err != nil {
		return err
	}
	*result = "ok"
	return nil
}

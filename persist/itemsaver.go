package persist

import (
	"context"
	"errors"
	"go-crawl/engine"
	"log"

	"github.com/olivere/elastic"
)

func ItemSaver(indexName string) (chan engine.Item, error) {

	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		count := 0
		for {
			item := <-out
			count++
			log.Printf("Item Saver : Got Item #%d : %v", count, item)
			err := Save(client, indexName, item)
			if err != nil {
				log.Printf("Item Saver : occure error : %v", err)
				continue
			}
		}
	}()
	return out, nil
}

func Save(client *elastic.Client, indexName string, item engine.Item) (err error) {

	if err != nil {
		return err
	}
	if item.Type == "" {
		return errors.New("must supply Type")
	}
	indexService := client.Index().
		Index(indexName).
		Type(item.Type).
		BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err = indexService.Do(context.Background())

	if err != nil {
		return err
	}

	return err

}

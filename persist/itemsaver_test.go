package persist

import (
	"go-crawl/engine"
	"go-crawl/model"
	"testing"

	"github.com/olivere/elastic"
)

func TestSave(t *testing.T) {

	var item engine.Item
	profile := model.Profile{
		Name:     "知心交友",
		Gender:   "男士",
		Age:      38,
		Height:   172,
		Income:   "12001-20000元",
		Marriage: "离异",
		House:    "内蒙古阿拉善盟",
	}
	item.Url = "http://l"
	item.Type = "zhenai"
	item.Id = "11104459"
	item.Payload = profile

	client, err := elastic.NewClient(
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}
	save(client, item)
}

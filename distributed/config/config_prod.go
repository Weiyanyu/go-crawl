// +build prod

package config

const (

	//Item service
	ItemSaverHost        = "go-crawl-itemserver"
	ItemSaverPort        = 6000
	ItemSaverESIndex     = "dating_profile"
	ItemSaverRpcFuncName = "ItemSaverService.Save"

	//worker service
	WorkerRpcFuncName = "CrawlService.Process"

	//Parser name
	ParseCityAndProfile = "ParseCityAndProfile"
	ParseCityList       = "ParseCityList"
	NilParser           = "NilParser"

	ElasticSearchUrl = "http://elasticsearch-server:9200"
)

var (
	WorkerHostList = []string{
		"go-crawl-workerserver:9000",
		":9001",
		":9002",
		":9003",
		":9004",
	}
)

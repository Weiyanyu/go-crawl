package config

const (

	//Item service
	ItemSaverPort        = 6000
	ItemSaverESIndex     = "dating_profile"
	ItemSaverRpcFuncName = "ItemSaverService.Save"

	//worker service
	WorkerRpcFuncName = "CrawlService.Process"

	//Parser name
	ParseCityAndProfile = "ParseCityAndProfile"
	ParseCityList       = "ParseCityList"
	NilParser           = "NilParser"
)

var (
	WorkerHostList = []string{
		":9000",
		":9001",
		":9002",
		":9003",
		":9004",
	}
)

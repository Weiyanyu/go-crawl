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

	//worker port
	WorkerPort0 = 9000
)

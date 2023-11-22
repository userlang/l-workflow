package config

const (
	NConfigAddress             = "172.16.0.51"
	NConfigPort                = 8848
	NConfigNamespaceId         = "dev" // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
	NConfigTimeoutMs           = 5000
	NConfigNotLoadCacheAtStart = true
	NConfigLogLevel            = "debug"
	NConfigDataId              = "workflow-dev.yml"
	NConfigGroup               = "DEV_GROUP"
	NServerAddress             = "172.16.0.51"
	NServerPort                = 8848
	NServerServiceName         = "workflow"
	NServerGroupName           = "DEV_GROUP"
	NServerWeight              = 1
	RedisAddressPort           = "172.16.0.166:22001"
	RedisPwd                   = "w3DEhqL67fF12"
	RedisDB                    = 0
	MySqlUsername              = "gbjk@test"
	MySqlPassword              = "gbjk@testpwd1"
	MySqlHost                  = "172.16.0.104"
	MySqlPort                  = "13049"
	MySqlDatabase              = "gbjk_workflow"
)

package config

const (
	NConfigAddress             = "***"
	NConfigPort                = 0
	NConfigNamespaceId         = "***" // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
	NConfigTimeoutMs           = 0
	NConfigNotLoadCacheAtStart = false
	NConfigLogLevel            = "***"
	NConfigDataId              = "***"
	NConfigGroup               = "***"
	NServerAddress             = "***"
	NServerPort                = 0
	NServerServiceName         = "***"
	NServerGroupName           = "***"
	NServerWeight              = 0
	RedisAddressPort           = "***"
	RedisPwd                   = "***"
	RedisDB                    = 0
	MySqlUsername              = "***"
	MySqlPassword              = "***"
	MySqlHost                  = "***"
	MySqlPort                  = 0
	MySqlDatabase              = "***"
)

package naocs

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"gopkg.in/yaml.v2"
	"log"
	"workflow/config"
)

var nval Nc

func Start() {

	// 创建serverConfig
	// 支持多个;至少一个ServerConfig
	serverConfig := []constant.ServerConfig{
		{
			IpAddr: config.NConfigAddress,
			Port:   config.NConfigPort,
		},
	}
	// 创建clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         config.NConfigNamespaceId, // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           config.NConfigTimeoutMs,
		NotLoadCacheAtStart: config.NConfigNotLoadCacheAtStart,
		LogLevel:            config.NConfigLogLevel,
	}

	nConfig(serverConfig, clientConfig)
	nServer(serverConfig, clientConfig)

}

func nServer(serverConfig []constant.ServerConfig, clientConfig constant.ClientConfig) {

	// 创建服务实例
	namingClient, err := clients.NewNamingClient(vo.NacosClientParam{
		ServerConfigs: serverConfig,
		ClientConfig:  &clientConfig,
	})

	if err != nil {
		log.Fatalf("创建NamingClient失败: %s", err.Error())
	}

	// 注册服务
	success, err := namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          config.NServerAddress,
		Port:        config.NServerPort,
		ServiceName: config.NServerServiceName,
		Weight:      config.NServerWeight,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		GroupName:   config.NServerGroupName,
		//Metadata: map[string]string{
		//	"groupName": "DEV_GROUP",
		//},
	})

	if err != nil {
		log.Fatalf("注册服务失败: %s", err.Error())
	}

	if success {
		fmt.Println("服务注册成功")
	} else {
		fmt.Println("服务注册失败")
	}

}

func nConfig(serverConfig []constant.ServerConfig, clientConfig constant.ClientConfig) {

	// 创建动态配置客户端
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfig,
		"clientConfig":  clientConfig,
	})

	if err != nil {
		log.Fatalf("初始化nacos失败: %s", err.Error())
	}

	// 获取配置
	dataId := config.NConfigDataId
	group := config.NConfigGroup
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group})

	if err != nil {
		log.Fatalf("获取%s配置失败: %s", dataId, err.Error())
	}

	err1 := yaml.Unmarshal([]byte(content), &nval)
	if err1 != nil {
		fmt.Println(err1)
	}

	fmt.Println(nval.Base.NConfigGroup, "正在测试读取配置信息")

	// 监听配置变化
	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("监听到变化")
			fmt.Println("group:" + group + ", dataId:" + dataId)
			fmt.Println("data: ", data)
		},
	})

	if err != nil {
		log.Fatalf("获取%s配置失败: %s", dataId, err.Error())
	}

}

type Nc struct {
	Base struct {
		NConfigAddress             string `yaml:"NConfigAddress"`
		NConfigPort                int    `yaml:"NConfigPort"`
		NConfigNamespaceId         string `yaml:"NConfigNamespaceId"`
		NConfigTimeoutMs           int    `yaml:"NConfigTimeoutMs"`
		NConfigNotLoadCacheAtStart string `yaml:"NConfigNotLoadCacheAtStart"`
		NConfigLogLevel            string `yaml:"NConfigLogLevel"`
		NConfigDataId              string `yaml:"NConfigDataId"`
		NConfigGroup               string `yaml:"NConfigGroup"`
		NServerAddress             string `yaml:"NServerAddress"`
		NServerPort                int    `yaml:"NServerPort"`
		NServerServiceName         string `yaml:"NServerServiceName"`
		NServerGroupName           string `yaml:"NServerGroupName"`
		NServerWeight              int    `yaml:"NServerWeight"`
	} `yaml:"base"`
}

package naocs

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
)

func Start() {

	// 创建serverConfig
	// 支持多个;至少一个ServerConfig
	serverConfig := []constant.ServerConfig{
		{
			IpAddr: "172.16.0.51",
			Port:   8848,
		},
	}

	// 创建clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         "dev", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogLevel:            "debug",
	}
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
		Ip:          "172.16.0.51",
		Port:        8848,
		ServiceName: "workflow",
		Weight:      1,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		GroupName:   "DEV_GROUP",
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

	// 创建动态配置客户端
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfig,
		"clientConfig":  clientConfig,
	})

	if err != nil {
		log.Fatalf("初始化nacos失败: %s", err.Error())
	}

	// 获取配置
	dataId := "workflow-dev.yml"
	group := "DEV_GROUP"
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group})

	if err != nil {
		log.Fatalf("获取%s配置失败: %s", dataId, err.Error())
	}

	fmt.Println(content)

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

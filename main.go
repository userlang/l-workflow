package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
	"workflow/naocs"
	"workflow/redis"
	"workflow/repository"
	"workflow/route"
)

func main() {

	config, err := ini.Load("./workflow.ini")
	if err != nil {
		fmt.Println("加载失败")
		os.Exit(1)
	}
	mysqlIp := config.Section("mysql").Key("ip").String()
	redisIp := config.Section("redis").Key("ip").String()
	appname := config.Section("").Key("appname").String()

	fmt.Println(mysqlIp)
	fmt.Println(redisIp)
	fmt.Println(appname)

	naocs.Start()
	repository.InitDB()
	redis.Start()
	route.Start()

}

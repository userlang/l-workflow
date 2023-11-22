package route

import (
	"github.com/gin-gonic/gin"
	"workflow/controllers"
)

func Start() {

	// 1 创建一个路由
	engine := gin.Default()
	//绑定路由规则 执行函数
	//gin.Context 相当于把Request  Response 封装起来

	engine.GET("/queryWorkFlowList", controllers.QueryWorkFlowListApi)
	engine.POST("/submit", controllers.SubmitApi)

	engine.Run(":8080")
}

package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"workflow/common"
	"workflow/repository/mysql"
)

func QueryWorkFlowListApi(context *gin.Context) {
	list := mysql.QueryWorkFlowList()
	res := common.ResponseData{Code: http.StatusOK, Data: list, Message: "success"}
	context.JSON(http.StatusOK, res)
}

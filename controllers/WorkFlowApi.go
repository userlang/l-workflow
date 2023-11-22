package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"workflow/common"
	"workflow/common/req"
	"workflow/repository/mysql"
)

/**查询审批流程列表*/
func QueryWorkFlowListApi(context *gin.Context) {
	list := mysql.QueryWorkFlowList()
	res := common.ResponseData{Code: http.StatusOK, Data: list, Message: "success"}
	context.JSON(res.Code, res)
}

/**提交审批流程*/
func SubmitApi(context *gin.Context) {
	var jsonObj req.JsonObj
	err := context.ShouldBindJSON(jsonObj)
	if err != nil {
		res := common.ResponseData{Code: http.StatusBadRequest, Data: nil, Message: "参数错误"}
		context.JSON(res.Code, res)
	}

}

/**审核通过*/
func Approved(context *gin.Context) {

}

/*审核拒绝*/
func Rejection(context *gin.Context) {

}

/**重新发起审批*/
func ReSubmit(context *gin.Context) {

}

/**查询当前业务审核信息详情*/
func QueryCurrentInstanceIdInfo(context *gin.Context) {

}

/**查询审批历史*/
func queryHistoryList(context *gin.Context) {

}

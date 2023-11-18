package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"workflow/models"
	"workflow/repository/mysql"
)

func Create(context *gin.Context) {

	instanceId := context.Param("instanceId")
	parseUint, err := strconv.ParseUint(instanceId, 10, 64)
	if err != nil {
		// 处理解析错误
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var approvalHistory = models.ApprovalHistory{Id: 111111, InstanceId: parseUint, StepId: 1, Participant: "22额", Result: "233", ResData: "2", BusinessId: 34, Reason: "reason", CreateTime: time.Now()}

	mysql.Create(&approvalHistory)

}

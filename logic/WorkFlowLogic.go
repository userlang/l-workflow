package logic

import (
	"fmt"
	"net/http"
	"time"
	"workflow/common"
	"workflow/common/req"
	"workflow/models"
	"workflow/repository/mysql"
)

func QueryWorkFlowListLogic() []models.WorkflowDefinition {
	return mysql.QueryWorkFlowList()
}

func Submit(obj req.JsonObj) common.ResponseData {
	/**查询审批流是否存在*/
	WorkflowDef := mysql.QueryInfoById(obj.WorkflowId)
	if WorkflowDef.Id == 0 {
		fmt.Println("审批流不存在")
		return common.ResponseData{Data: nil, Code: http.StatusOK, Message: "审批流不存在"}
	}
	/**查询审批流第一步骤*/
	StepDef := mysql.QueryStepDefinitionInfo(WorkflowDef.Id, 1)
	if StepDef.Id == 0 {
		fmt.Println("审批步骤为空")
		return common.ResponseData{Data: nil, Code: http.StatusOK, Message: "审批步骤为空"}

	}
	/**写入信息到流程实例表*/
	var ins models.WorkflowInstance
	ins.WorkflowDefinitionId = WorkflowDef.Id
	ins.CurrentStepId = StepDef.Id
	ins.BusinessId = obj.BusinessId
	ins.Status = common.Submit
	ins.Assignee = StepDef.Assignee
	ins.CreationTime = time.Now()
	ins.Creator = obj.Creator
	mysql.InsCreate(&ins)
	/**写入日志*/
	var history models.ApprovalHistory
	history.BusinessId = obj.BusinessId
	history.InstanceId = ins.Id
	history.StepId = 0
	history.Result = common.HisSubmit
	history.Participant = obj.Creator
	history.ResData = obj.BusinessData
	history.CreateTime = time.Now()
	mysql.HisCreate(&history)
	var data = make(map[string]int)
	data["instanceId"] = ins.Id
	data["workflowId"] = obj.WorkflowId
	data["businessId"] = obj.BusinessId
	return common.ResponseData{Data: data, Code: http.StatusOK, Message: "操作成功"}
}
func Approved(obj req.JsonObj) common.ResponseData {
	insInfo := mysql.QueryInsById(obj.InstanceId)
	if insInfo == nil {
		return common.ResponseData{Data: nil, Code: http.StatusOK, Message: "审批实例不存在"}
	}
	if insInfo.Status == common.Complete {
		return common.ResponseData{Data: nil, Code: http.StatusOK, Message: "已经审批完了不能再审批了"}
	}
	end, nextAssignee, nextNumber := GetNextStep(insInfo.Id)
	//记录审批日志信息
	var history models.ApprovalHistory
	history.BusinessId = obj.BusinessId
	history.InstanceId = insInfo.Id
	history.StepId = insInfo.CurrentStepId
	history.Result = common.Approved
	history.Participant = obj.Assignee
	history.ResData = obj.BusinessData
	history.CreateTime = time.Now()
	mysql.HisCreate(&history)
	//更新实例
	var updateIns models.WorkflowInstance
	if end {
		updateIns.Status = common.Complete
	} else {
		updateIns.Status = common.Approved
		updateIns.CurrentStepId = nextNumber
		updateIns.Assignee = nextAssignee
	}
	mysql.UpdateById(&updateIns)

	return common.ResponseData{Data: nil, Code: http.StatusOK, Message: "审核通过"}
}

func Rejection(obj req.JsonObj) common.ResponseData {

	return common.ResponseData{Data: nil, Code: http.StatusOK, Message: "驳回成功"}
}

func ReSubmit(obj req.JsonObj) common.ResponseData {

	return common.ResponseData{Data: nil, Code: http.StatusOK, Message: "重新发起成功"}
}

func QueryCurrentInstanceInfo(obj req.JsonObj) common.ResponseData {

	return common.ResponseData{Data: nil, Code: http.StatusOK, Message: "查询成功"}
}
func QueryHistoryList(obj req.JsonObj) common.ResponseData {

	return common.ResponseData{Data: nil, Code: http.StatusOK, Message: "查询成功"}
}

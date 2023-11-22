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
	end, nextAssignee, nextNumber := GetNextStep(insInfo.Id)
	//更新实例
	var updateIns models.WorkflowInstance
	updateIns.Id = insInfo.Id
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

	insInfo := mysql.QueryInsById(obj.InstanceId)
	if insInfo == nil {
		return common.ResponseData{Data: nil, Code: http.StatusOK, Message: "没有审批实例"}
	}
	if insInfo.Status == common.Complete {
		return common.ResponseData{Data: nil, Code: http.StatusOK, Message: "流程已经审核完毕 无法拒绝"}
	}
	if insInfo.Status == common.Submit {
		return common.ResponseData{Data: nil, Code: http.StatusOK, Message: "流程节点是开始节点 无法拒绝"}
	}
	//记录审批日志信息
	var history models.ApprovalHistory
	history.BusinessId = obj.BusinessId
	history.InstanceId = insInfo.Id
	history.StepId = insInfo.CurrentStepId
	history.Result = common.HisRefuse
	history.Participant = obj.Assignee
	history.Reason = obj.Reason
	history.ResData = obj.BusinessData
	history.CreateTime = time.Now()
	mysql.HisCreate(&history)
	//更新上一步实例
	start, preAssignee, preNumber := GetPreStep(insInfo.Id)
	//更新实例
	var updateIns models.WorkflowInstance
	updateIns.Id = insInfo.Id
	if !start {
		updateIns.CurrentStepId = preNumber
		updateIns.Assignee = preAssignee
	}
	updateIns.Status = common.Refuse

	mysql.UpdateById(&updateIns)

	return common.ResponseData{Data: nil, Code: http.StatusOK, Message: "驳回成功"}
}

func ReSubmit(obj req.JsonObj) common.ResponseData {

	insInfo := mysql.QueryInsById(obj.InstanceId)
	if insInfo == nil {
		return common.ResponseData{Data: nil, Code: http.StatusOK, Message: "没有审批实例"}
	}
	if insInfo.Status != common.Refuse {
		return common.ResponseData{Data: nil, Code: http.StatusOK, Message: "只有拒绝了才可以重新提交哦"}
	}
	var updateIns models.WorkflowInstance

	updateIns.Status = common.Submit
	updateIns.Id = insInfo.Id

	mysql.UpdateById(&updateIns)

	return common.ResponseData{Data: nil, Code: http.StatusOK, Message: "重新发起成功"}
}

func QueryCurrentInstanceInfo(obj req.JsonObj) common.ResponseData {
	InsInfo := mysql.QueryInsById(obj.InstanceId)
	return common.ResponseData{Data: InsInfo, Code: http.StatusOK, Message: "查询成功"}
}
func QueryHistoryList(obj req.JsonObj) common.ResponseData {
	list := mysql.QueryHisByBusIdList(obj.BusinessId)
	return common.ResponseData{Data: list, Code: http.StatusOK, Message: "查询成功"}
}

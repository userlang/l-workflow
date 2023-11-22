package controllers

import (
	"fmt"
	"workflow/common"
	"workflow/common/req"
	"workflow/models"
	"workflow/repository/mysql"
)

func queryWorkFlowListLogic() []models.WorkflowDefinition {
	return mysql.QueryWorkFlowList()
}

func submit(obj req.JsonObj) bool {
	/**查询审批流是否存在*/
	WorkflowDef := mysql.QueryInfoById(obj.WorkflowId)
	if WorkflowDef.Id == 0 {
		fmt.Println("审批流不存在")
		return false
	}
	/**查询审批流第一步骤*/
	StepDef := mysql.QueryStepDefinitionInfo(WorkflowDef.Id, 1)
	if StepDef.Id == 0 {
		fmt.Println("审批步骤为空")
		return false
	}
	/**写入信息到流程实例表*/
	var ins models.WorkflowInstance
	ins.WorkflowDefinitionId = WorkflowDef.Id
	ins.CurrentStepId = StepDef.Id
	ins.BusinessId = obj.BusinessId
	ins.Status = common.Submit
	ins.Assignee = StepDef.Assignee
	ins.Creator = obj.Creator
	mysql.InsCreate(&ins)
	/**写入日志*/
	var history models.ApprovalHistory
	history.BusinessId = obj.BusinessId
	history.InstanceId = ins.Id

	return false
}

package logic

import "workflow/repository/mysql"

func GetPreStep(instanceId int) (start bool, preAssignee string, preNumber int) {

	// 省略
	return false, "", 1
}

func GetNextStep(instanceId int) (end bool, nextAssignee string, nextNumber int) {
	insInfo := mysql.QueryInsById(instanceId)
	stepInfo := mysql.QueryStepDefinitionById(insInfo.CurrentStepId)

	if stepInfo.NextStepId == -1 {
		end = true
	} else {
		nextStep := mysql.QueryStepDefinitionById(stepInfo.NextStepId)
		nextAssignee = nextStep.Assignee
		nextNumber = nextStep.Id
		end = false
	}
	return end, nextAssignee, nextNumber
}

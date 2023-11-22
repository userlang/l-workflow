package logic

import "workflow/repository/mysql"

func GetPreStep(instanceId int) (start bool, preAssignee string, preNumber int) {
	insInfo := mysql.QueryInsById(instanceId)
	stepInfo := mysql.QueryStepDefinitionById(insInfo.CurrentStepId)
	if stepInfo.PreStepId == -1 {
		start = true
	} else {
		preStep := mysql.QueryStepDefinitionById(stepInfo.PreStepId)
		preAssignee = preStep.Assignee
		preNumber = preStep.Id
		start = false
	}
	// 省略
	return start, preAssignee, preNumber
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

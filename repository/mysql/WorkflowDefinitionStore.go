package mysql

import (
	"workflow/models"
	"workflow/repository"
)

func QueryWorkFlowList() []models.WorkflowDefinition {

	var list []models.WorkflowDefinition
	repository.DB.Find(&list)
	return list
}

func QueryInfoById(id int) *models.WorkflowDefinition {
	var info models.WorkflowDefinition
	repository.DB.First(&info, id)
	return &info
}

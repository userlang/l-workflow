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

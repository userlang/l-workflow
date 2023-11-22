package mysql

import (
	"workflow/models"
	"workflow/repository"
)

func QueryStepDefinitionInfo(definitionId int, stepNumber int) *models.WorkflowStepDefinition {
	var info models.WorkflowStepDefinition
	result := repository.DB.Where("workflow_definition_id=? and step_number=? ", definitionId, stepNumber).First(&info)
	if result.Error != nil {
		return nil
	}

	return &info
}

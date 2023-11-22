package mysql

import (
	"workflow/models"
	"workflow/repository"
)

func InsCreate(info *models.WorkflowInstance) {
	repository.DB.Create(info)
}

func QueryInsById(id int) *models.WorkflowInstance {
	var info models.WorkflowInstance
	repository.DB.First(&info, id)
	return &info
}

func UpdateById(info *models.WorkflowInstance) {
	updateFields := make(map[string]interface{})
	if info.CurrentStepId != 0 {
		updateFields["current_step_id"] = info.CurrentStepId
	}
	if info.Status != "0" {
		updateFields["status"] = info.Status
	}
	if info.Assignee != "" {
		updateFields["assignee"] = info.Assignee
	}

	repository.DB.Model(&models.WorkflowDefinition{}).Where("id= ? ", info.Id).Updates(updateFields)
}

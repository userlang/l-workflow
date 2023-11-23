package mysql

import (
	"gorm.io/gorm"
	"workflow/models"
	"workflow/repository"
)

func InsCreate(info *models.WorkflowInstance, tx *gorm.DB) {

	if tx != nil {
		err := tx.Create(info).Error
		if err != nil {
			tx.Rollback()
		}
	} else {
		repository.DB.Create(info)
	}

}

func QueryInsById(id int) *models.WorkflowInstance {
	var info models.WorkflowInstance
	repository.DB.First(&info, id)
	return &info
}

func UpdateById(info *models.WorkflowInstance, tx *gorm.DB) {
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
	if tx != nil {
		err := tx.Model(&models.WorkflowInstance{}).Where("id= ? ", info.Id).Updates(updateFields).Error
		if err != nil {
			tx.Rollback()
		}
	} else {
		repository.DB.Model(&models.WorkflowInstance{}).Where("id= ? ", info.Id).Updates(updateFields)
	}

}

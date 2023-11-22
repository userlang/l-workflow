package mysql

import (
	"workflow/models"
	"workflow/repository"
)

func InsCreate(info *models.WorkflowInstance) {
	repository.DB.Create(info)
}

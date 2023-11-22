package mysql

import (
	"workflow/models"
	"workflow/repository"
)

func HisCreate(history *models.ApprovalHistory) {

	repository.DB.Create(history)

}

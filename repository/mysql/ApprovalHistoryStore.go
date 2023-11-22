package mysql

import (
	"workflow/models"
	"workflow/repository"
)

func HisCreate(history *models.ApprovalHistory) {

	repository.DB.Create(history)

}

func QueryHisByBusIdList(busId int) *models.ApprovalHistory {

	var info models.ApprovalHistory
	repository.DB.Where("business_id", busId).First(&info)
	return &info
}

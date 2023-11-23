package mysql

import (
	"gorm.io/gorm"
	"workflow/models"
	"workflow/repository"
)

func HisCreate(history *models.ApprovalHistory, tx *gorm.DB) {

	if tx != nil {
		err := tx.Create(history).Error
		if err != nil {
			tx.Rollback()
		}
	} else {
		repository.DB.Create(history)
	}

}

func QueryHisByBusIdList(busId int) *models.ApprovalHistory {

	var info models.ApprovalHistory
	repository.DB.Where("business_id", busId).First(&info)
	return &info
}

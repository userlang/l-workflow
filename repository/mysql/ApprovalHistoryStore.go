package mysql

import (
	"workflow/models"
	"workflow/repository"
)

func CreateApprovalHistory(approvalhistory *models.ApprovalHistory) {
	//db, err := repository.InitDB()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//db.Create(approvalhistory)
	repository.DB.Create(approvalhistory)
}

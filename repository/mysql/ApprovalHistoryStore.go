package mysql

import (
	"fmt"
	"workflow/models"
	"workflow/repository"
)

func Create(approvalhistory *models.ApprovalHistory) {
	db, err := repository.InitDB()
	if err != nil {
		fmt.Println(err)
	}
	db.Create(approvalhistory)
}

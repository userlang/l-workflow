package controllers

import (
	"workflow/models"
	"workflow/repository/mysql"
)

func queryWorkFlowListLogic() []models.WorkflowDefinition {
	return mysql.QueryWorkFlowList()
}

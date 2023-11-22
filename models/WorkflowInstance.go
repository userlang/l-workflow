package models

import "time"

type WorkflowInstance struct {
	Id                   int       `gorm:"column:id" json:"id"`
	WorkflowDefinitionId int       `gorm:"column:workflow_definition_id" json:"workflow_definition_id"`
	CurrentStepId        int       `gorm:"column:current_step_id" json:"current_step_id"`
	BusinessId           int       `gorm:"column:business_id" json:"business_id"`
	Status               string    `gorm:"column:status" json:"status"`
	CreationTime         time.Time `gorm:"column:creation_time" json:"creation_time"`
	Assignee             string    `gorm:"column:assignee" json:"assignee"`
	Creator              string    `gorm:"column:creator" json:"creator"`
}

func (WorkflowInstance) TableName() string {

	return "workflow_instance"
}

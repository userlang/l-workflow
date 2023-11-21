package models

import "time"

type WorkflowInstance struct {
	Id                   uint64    `gorm:"column:id" json:"id"`
	WorkflowDefinitionId uint64    `gorm:"column:workflow_definition_id" json:"workflow_definition_id"`
	CurrentStepId        uint64    `gorm:"column:current_step_id" json:"current_step_id"`
	BusinessId           uint64    `gorm:"column:business_id" json:"business_id"`
	Status               string    `gorm:"column:status" json:"status"`
	CreationTime         time.Time `gorm:"column:creation_time" json:"creation_time"`
	Assignee             string    `gorm:"column:assignee" json:"assignee"`
	Creator              string    `gorm:"column:creator" json:"creator"`
}

func (WorkflowInstance) TableName() string {

	return "workflow_instance"
}

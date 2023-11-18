package models

import "time"

type workflowInstance struct {
	Id                   uint64    `gorm:"column:id"`
	WorkflowDefinitionId uint64    `gorm:"column:workflow_definition_id"`
	CurrentStepId        uint64    `gorm:"column:current_step_id"`
	BusinessId           uint64    `gorm:"column:business_id"`
	Status               string    `gorm:"column:status"`
	CreationTime         time.Time `gorm:"column:creation_time"`
	assignee             string    `gorm:"column:assignee"`
	creator              string    `gorm:"column:creator"`
}

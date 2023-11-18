package models

import "time"

type workflowTask struct {
	Id                 uint64    `gorm:"column:id;primaryKey"`
	WorkflowInstanceId uint64    `gorm:"column:workflow_instance_id"`
	StepDefinitionId   uint64    `gorm:"column:step_definition_id"`
	BusinessId         uint64    `gorm:"column:business_id"`
	Assignee           string    `gorm:"column:assignee"`
	Status             string    `gorm:"column:status"`
	CreationTime       time.Time `gorm:"column:creation_time"`
	CompletionTime     time.Time `gorm:"column:completion_time"`
}

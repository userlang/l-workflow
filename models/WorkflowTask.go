package models

import "time"

type WorkflowTask struct {
	Id                 uint64    `gorm:"column:id;primaryKey" json:"id"`
	WorkflowInstanceId uint64    `gorm:"column:workflow_instance_id" json:"workflow_instance_id"`
	StepDefinitionId   uint64    `gorm:"column:step_definition_id" json:"step_definition_id"`
	BusinessId         uint64    `gorm:"column:business_id" json:"business_id"`
	Assignee           string    `gorm:"column:assignee" json:"assignee"`
	Status             string    `gorm:"column:status" json:"status"`
	CreationTime       time.Time `gorm:"column:creation_time" json:"creation_time"`
	CompletionTime     time.Time `gorm:"column:completion_time" json:"completion_time"`
}

func (WorkflowTask) TableName() string {

	return "workflow_task"
}

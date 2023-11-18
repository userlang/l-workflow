package models

type workflowStepDefinition struct {
	Id                   uint64 `gorm:"column:id;primaryKey"`
	WorkflowDefinitionId uint64 `gorm:"column:workflow_definition_id"`
	StepNumber           uint64 `gorm:"column:step_number"`
	StepName             uint64 `gorm:"column:step_name"`
	Assignee             uint64 `gorm:"column:assignee"`
	NextStepId           uint64 `gorm:"column:next_step_id"`
	CreateTime           uint64 `gorm:"column:create_time"`
	PreStepId            uint64 `gorm:"column:pre_step_id"`
}

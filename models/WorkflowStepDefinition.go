package models

type WorkflowStepDefinition struct {
	Id                   uint64 `gorm:"column:id;primaryKey" json:"id"`
	WorkflowDefinitionId uint64 `gorm:"column:workflow_definition_id" json:"workflow_definition_id"`
	StepNumber           uint64 `gorm:"column:step_number" json:"step_number"`
	StepName             uint64 `gorm:"column:step_name" json:"step_name"`
	Assignee             uint64 `gorm:"column:assignee" json:"assignee"`
	NextStepId           uint64 `gorm:"column:next_step_id" json:"next_step_id"`
	CreateTime           uint64 `gorm:"column:create_time" json:"create_time"`
	PreStepId            uint64 `gorm:"column:pre_step_id" json:"pre_step_id"`
}

func (WorkflowStepDefinition) TableName() string {

	return "workflow_step_definition"
}

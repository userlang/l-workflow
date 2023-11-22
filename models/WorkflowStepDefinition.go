package models

type WorkflowStepDefinition struct {
	Id                   int    `gorm:"column:id;primaryKey" json:"id"`
	WorkflowDefinitionId int    `gorm:"column:workflow_definition_id" json:"workflow_definition_id"`
	StepNumber           int    `gorm:"column:step_number" json:"step_number"`
	StepName             int    `gorm:"column:step_name" json:"step_name"`
	Assignee             string `gorm:"column:assignee" json:"assignee"`
	NextStepId           int    `gorm:"column:next_step_id" json:"next_step_id"`
	CreateTime           int    `gorm:"column:create_time" json:"create_time"`
	PreStepId            int    `gorm:"column:pre_step_id" json:"pre_step_id"`
}

func (WorkflowStepDefinition) TableName() string {

	return "workflow_step_definition"
}

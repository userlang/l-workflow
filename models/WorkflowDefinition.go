package models

import "time"

type WorkflowDefinition struct {
	Id          int       `gorm:"column:id;primaryKey" json:"id"`
	Title       string    `gorm:"column:title" json:"title"`
	Description string    `gorm:"column:description" json:"description"`
	CreateTime  time.Time `gorm:"column:create_time" json:"create_time"`
}

func (WorkflowDefinition) TableName() string {

	return "workflow_definition"
}

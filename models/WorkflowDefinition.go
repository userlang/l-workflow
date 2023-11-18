package models

import "time"

type WorkflowDefinition struct {
	Id          uint64    `gorm:"column:id;primaryKey"`
	Title       string    `gorm:"column:title"`
	Description string    `gorm:"column:description"`
	CreateTime  time.Time `gorm:"column:create_time"`
}

func (WorkflowDefinition) TableName() string {

	return "workflow_definition"
}

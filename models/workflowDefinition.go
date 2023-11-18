package models

import "time"

type workflowDefinition struct {
	Id          uint64    `gorm:"column:id;primaryKey"`
	Title       string    `gorm:"column:title"`
	Description string    `gorm:"column:description"`
	CreateTime  time.Time `gorm:"column:create_time"`
}

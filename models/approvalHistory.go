package models

import "time"

type approvalHistory struct {
	Id          uint64    `gorm:"column:id;primaryKey"`
	InstanceId  uint64    `gorm:"column:instance_id"`
	StepId      uint64    `gorm:"column:step_id"`
	Participant string    `gorm:"column:participant"`
	Result      string    `gorm:"column:result"`
	ResData     string    `gorm:"column:res_data"`
	BusinessId  uint64    `gorm:"column:business_id"`
	Reason      string    `gorm:"column:reason"`
	CreateTime  time.Time `gorm:"create_time"`
}

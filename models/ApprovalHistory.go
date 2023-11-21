package models

import "time"

type ApprovalHistory struct {
	Id          uint64    `gorm:"column:id;primaryKey" json:"id"`
	InstanceId  uint64    `gorm:"column:instance_id" json:"instance_id"`
	StepId      uint64    `gorm:"column:step_id" json:"step_id"`
	Participant string    `gorm:"column:participant" json:"participant"`
	Result      string    `gorm:"column:result" json:"result"`
	ResData     string    `gorm:"column:res_data" json:"res_data"`
	BusinessId  uint64    `gorm:"column:business_id" json:"business_id"`
	Reason      string    `gorm:"column:reason" json:"reason"`
	CreateTime  time.Time `gorm:"create_time" json:"create_time"`
}

func (ApprovalHistory) TableName() string {

	return "approval_history"
}

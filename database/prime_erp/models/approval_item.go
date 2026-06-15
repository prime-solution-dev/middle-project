package models

import (
	"time"
	"gorm.io/datatypes"
)

type ApprovalItem struct {
	Id string `gorm:"column:id;primaryKey" json:"id"`
	ApprovalId string `gorm:"column:approval_id" json:"approval_id"`
	StepSeq int32 `gorm:"column:step_seq" json:"step_seq"`
	IsCondition bool `gorm:"column:is_condition" json:"is_condition"`
	Condition datatypes.JSON `gorm:"column:condition" json:"condition"`
	Status string `gorm:"column:status" json:"status"`
	ActionBy string `gorm:"column:action_by" json:"action_by"`
	ActionDate time.Time `gorm:"column:action_date" json:"action_date"`
	CreateDate time.Time `gorm:"column:create_date" json:"create_date"`
	CreateBy string `gorm:"column:create_by" json:"create_by"`
	UpdateDate time.Time `gorm:"column:update_date" json:"update_date"`
	UpdateBy string `gorm:"column:update_by" json:"update_by"`
}

func (ApprovalItem) TableName() string {
	return "approval_item"
}

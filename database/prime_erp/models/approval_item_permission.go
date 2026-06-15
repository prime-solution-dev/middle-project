package models

import (
	"time"
	"gorm.io/datatypes"
)

type ApprovalItemPermission struct {
	Id string `gorm:"column:id;primaryKey" json:"id"`
	ApprovalItemId string `gorm:"column:approval_item_id" json:"approval_item_id"`
	UserCode string `gorm:"column:user_code" json:"user_code"`
}

func (ApprovalItemPermission) TableName() string {
	return "approval_item_permission"
}

package models

import (
	"time"
	"gorm.io/datatypes"
)

type Approval struct {
	Id string `gorm:"column:id" json:"id"`
	ApproveCode string `gorm:"column:approve_code" json:"approve_code"`
}

func (Approval) TableName() string {
	return "approval"
}

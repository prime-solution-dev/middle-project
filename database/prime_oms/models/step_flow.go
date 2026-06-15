package models

import (
	"time"
	"gorm.io/datatypes"
)

type StepFlow struct {
	Id string `gorm:"column:id" json:"id"`
	SrcStepCode string `gorm:"column:src_step_code" json:"src_step_code"`
	SrcStatus string `gorm:"column:src_status" json:"src_status"`
	DestStepCode string `gorm:"column:dest_step_code" json:"dest_step_code"`
	DestStatus string `gorm:"column:dest_status" json:"dest_status"`
}

func (StepFlow) TableName() string {
	return "step_flow"
}

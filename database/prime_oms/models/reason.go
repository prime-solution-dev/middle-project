package models

import (
	"time"
	"gorm.io/datatypes"
)

type Reason struct {
	Id string `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	Code string `gorm:"column:code" json:"code"`
	ReasonType string `gorm:"column:reason_type" json:"reason_type"`
	CreateBy string `gorm:"column:create_by" json:"create_by"`
	Seq int32 `gorm:"column:seq" json:"seq"`
	CreateDtm time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateBy string `gorm:"column:update_by" json:"update_by"`
	UpdateDtm time.Time `gorm:"column:update_dtm" json:"update_dtm"`
}

func (Reason) TableName() string {
	return "reason"
}

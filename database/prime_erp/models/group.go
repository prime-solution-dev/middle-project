package models

import (
	"time"
	"gorm.io/datatypes"
)

type Group struct {
	Id string `gorm:"column:id;primaryKey" json:"id"`
	GroupCode string `gorm:"column:group_code" json:"group_code"`
	GroupName string `gorm:"column:group_name" json:"group_name"`
	Value string `gorm:"column:value" json:"value"`
	ValueInt float64 `gorm:"column:value_int" json:"value_int"`
	Seq int32 `gorm:"column:seq" json:"seq"`
	CreateDtm time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	CreateBy string `gorm:"column:create_by" json:"create_by"`
	UpdateDtm time.Time `gorm:"column:update_dtm" json:"update_dtm"`
	UpdateBy string `gorm:"column:update_by" json:"update_by"`
}

func (Group) TableName() string {
	return "group"
}

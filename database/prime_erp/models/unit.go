package models

import (
	"time"
	"gorm.io/datatypes"
)

type Unit struct {
	Id string `gorm:"column:id;primaryKey" json:"id"`
	Topic string `gorm:"column:topic" json:"topic"`
	UnitCode string `gorm:"column:unit_code" json:"unit_code"`
	UnitName string `gorm:"column:unit_name" json:"unit_name"`
}

func (Unit) TableName() string {
	return "unit"
}

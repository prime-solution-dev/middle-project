package models

import (
	"time"
	"gorm.io/datatypes"
)

type Time struct {
	Id string `gorm:"column:id;primaryKey" json:"id"`
	Topic string `gorm:"column:topic" json:"topic"`
	Code string `gorm:"column:code" json:"code"`
	Name string `gorm:"column:name" json:"name"`
	StartTime time.Time `gorm:"column:start_time" json:"start_time"`
	EndTime time.Time `gorm:"column:end_time" json:"end_time"`
}

func (Time) TableName() string {
	return "time"
}

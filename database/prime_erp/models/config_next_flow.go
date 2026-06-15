package models

import (
	"time"
	"gorm.io/datatypes"
)

type ConfigNextFlow struct {
	Module string `gorm:"column:module" json:"module"`
	Type string `gorm:"column:type" json:"type"`
	IsNextFlow bool `gorm:"column:is_next_flow" json:"is_next_flow"`
	Topic string `gorm:"column:topic" json:"topic"`
}

func (ConfigNextFlow) TableName() string {
	return "config_next_flow"
}

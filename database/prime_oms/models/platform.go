package models

import (
	"time"
	"gorm.io/datatypes"
)

type Platform struct {
	Id string `gorm:"column:id" json:"id"`
	PlatformCode string `gorm:"column:platform_code" json:"platform_code"`
	PlatformName string `gorm:"column:platform_name" json:"platform_name"`
	Active bool `gorm:"column:active" json:"active"`
	PlatformType string `gorm:"column:platform_type" json:"platform_type"`
	PlatformShortName string `gorm:"column:platform_short_name" json:"platform_short_name"`
}

func (Platform) TableName() string {
	return "platform"
}

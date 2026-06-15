package models

import (
	"time"
	"gorm.io/datatypes"
)

type PlatformLog struct {
	Id int32 `gorm:"column:id" json:"id"`
	Platform string `gorm:"column:platform" json:"platform"`
	Action string `gorm:"column:action" json:"action"`
	Request datatypes.JSON `gorm:"column:request" json:"request"`
	Response datatypes.JSON `gorm:"column:response" json:"response"`
	Status string `gorm:"column:status" json:"status"`
	ErrorMessage string `gorm:"column:error_message" json:"error_message"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	ShopCode string `gorm:"column:shop_code" json:"shop_code"`
}

func (PlatformLog) TableName() string {
	return "platform_log"
}

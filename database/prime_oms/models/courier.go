package models

import (
	"time"
)

type Courier struct {
	Id          string    `gorm:"column:id" json:"id"`
	CourierType string    `gorm:"column:courier_type" json:"courier_type"`
	CourierCode string    `gorm:"column:courier_code" json:"courier_code"`
	CourierName string    `gorm:"column:courier_name" json:"courier_name"`
	CreateBy    string    `gorm:"column:create_by" json:"create_by"`
	CreateDtm   time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateBy    string    `gorm:"column:update_by" json:"update_by"`
	UpdateDtm   time.Time `gorm:"column:update_dtm" json:"update_dtm"`
	Name        string    `gorm:"column:name" json:"name"`
	Code        string    `gorm:"column:code" json:"code"`
	AuthType    string    `gorm:"column:auth_type" json:"auth_type"`
	Credentials string    `gorm:"column:credentials" json:"credentials"`
	IsActive    bool      `gorm:"column:is_active" json:"is_active"`
	IsReturn    bool      `gorm:"column:is_return" json:"is_return"`
}

func (Courier) TableName() string {
	return "courier"
}

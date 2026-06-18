package models

import (
	"time"
)

type ShopShipBill struct {
	Id            string    `gorm:"column:id" json:"id"`
	InterfaceCode string    `gorm:"column:interface_code" json:"interface_code"`
	OmsCode       string    `gorm:"column:oms_code" json:"oms_code"`
	Name          string    `gorm:"column:name" json:"name"`
	Type          string    `gorm:"column:type" json:"type"`
	CreateBy      string    `gorm:"column:create_by" json:"create_by"`
	CreateDtm     time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateBy      string    `gorm:"column:update_by" json:"update_by"`
	UpdateDtm     time.Time `gorm:"column:update_dtm" json:"update_dtm"`
	ShopId        string    `gorm:"column:shop_id" json:"shop_id"`
}

func (ShopShipBill) TableName() string {
	return "shop_ship_bill"
}

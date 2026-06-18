package models

import (
	"time"
)

type TradeDiscount struct {
	Id            string    `gorm:"column:id" json:"id"`
	DiscountCode  string    `gorm:"column:discount_code" json:"discount_code"`
	DiscountName  string    `gorm:"column:discount_name" json:"discount_name"`
	Type          string    `gorm:"column:type" json:"type"`
	CustomerCode  string    `gorm:"column:customer_code" json:"customer_code"`
	CustomerGroup string    `gorm:"column:customer_group" json:"customer_group"`
	StartDtm      time.Time `gorm:"column:start_dtm" json:"start_dtm"`
	EndDtm        time.Time `gorm:"column:end_dtm" json:"end_dtm"`
	QuotaType     string    `gorm:"column:quota_type" json:"quota_type"`
	QuotaOrder    int32     `gorm:"column:quota_order" json:"quota_order"`
	Status        string    `gorm:"column:status" json:"status"`
	CreateDtm     time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateDtm     time.Time `gorm:"column:update_dtm" json:"update_dtm"`
	CreateBy      string    `gorm:"column:create_by" json:"create_by"`
	UpdateBy      string    `gorm:"column:update_by" json:"update_by"`
}

func (TradeDiscount) TableName() string {
	return "trade_discount"
}

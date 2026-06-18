package models

import (
	"time"
)

type TradeDiscountItem struct {
	Id              string    `gorm:"column:id" json:"id"`
	TradeDiscountId string    `gorm:"column:trade_discount_id" json:"trade_discount_id"`
	Seq             int32     `gorm:"column:seq" json:"seq"`
	BillAmountStart float64   `gorm:"column:bill_amount_start" json:"bill_amount_start"`
	BillAmountEnd   float64   `gorm:"column:bill_amount_end" json:"bill_amount_end"`
	DiscountPercent float64   `gorm:"column:discount_percent" json:"discount_percent"`
	DiscountAmount  float64   `gorm:"column:discount_amount" json:"discount_amount"`
	CreateDtm       time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateDtm       time.Time `gorm:"column:update_dtm" json:"update_dtm"`
}

func (TradeDiscountItem) TableName() string {
	return "trade_discount_item"
}

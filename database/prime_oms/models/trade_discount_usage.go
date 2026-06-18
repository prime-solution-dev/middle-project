package models

import (
	"time"
)

type TradeDiscountUsage struct {
	Id                  string    `gorm:"column:id" json:"id"`
	TradeDiscountId     string    `gorm:"column:trade_discount_id" json:"trade_discount_id"`
	TradeDiscountCode   string    `gorm:"column:trade_discount_code" json:"trade_discount_code"`
	TradeDiscountItemId string    `gorm:"column:trade_discount_item_id" json:"trade_discount_item_id"`
	OrderNo             string    `gorm:"column:order_no" json:"order_no"`
	SalesOrderNo        string    `gorm:"column:sales_order_no" json:"sales_order_no"`
	CustomerCode        string    `gorm:"column:customer_code" json:"customer_code"`
	BillAmount          float64   `gorm:"column:bill_amount" json:"bill_amount"`
	DiscountValue       float64   `gorm:"column:discount_value" json:"discount_value"`
	FinalAmount         float64   `gorm:"column:final_amount" json:"final_amount"`
	TotalQty            int32     `gorm:"column:total_qty" json:"total_qty"`
	Status              string    `gorm:"column:status" json:"status"`
	CreateDtm           time.Time `gorm:"column:create_dtm" json:"create_dtm"`
}

func (TradeDiscountUsage) TableName() string {
	return "trade_discount_usage"
}

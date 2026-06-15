package models

import (
	"time"
	"gorm.io/datatypes"
)

type PlatformStatementItem struct {
	Id string `gorm:"column:id" json:"id"`
	DocumentId string `gorm:"column:document_id" json:"document_id"`
	PlatformOrderId string `gorm:"column:platform_order_id" json:"platform_order_id"`
	ProductId string `gorm:"column:product_id" json:"product_id"`
	ProductName string `gorm:"column:product_name" json:"product_name"`
	Qty float64 `gorm:"column:qty" json:"qty"`
	UnitPrice float64 `gorm:"column:unit_price" json:"unit_price"`
	SubtotalAmount float64 `gorm:"column:subtotal_amount" json:"subtotal_amount"`
	ProductCommissionFee float64 `gorm:"column:product_commission_fee" json:"product_commission_fee"`
	ProductPlatformFee float64 `gorm:"column:product_platform_fee" json:"product_platform_fee"`
	ProductSettlementAmount float64 `gorm:"column:product_settlement_amount" json:"product_settlement_amount"`
	Currency string `gorm:"column:currency" json:"currency"`
	CreateDtm time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateDtm time.Time `gorm:"column:update_dtm" json:"update_dtm"`
}

func (PlatformStatementItem) TableName() string {
	return "platform_statement_item"
}

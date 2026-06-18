package models

import (
	"time"
)

type PlatformStatementDocument struct {
	Id               string    `gorm:"column:id" json:"id"`
	StatementId      string    `gorm:"column:statement_id" json:"statement_id"`
	PaymentMethod    string    `gorm:"column:payment_method" json:"payment_method"`
	PaymentSubMethod string    `gorm:"column:payment_sub_method" json:"payment_sub_method"`
	DocumentRef      string    `gorm:"column:document_ref" json:"document_ref"`
	ShippingFee      float64   `gorm:"column:shipping_fee" json:"shipping_fee"`
	CommissionFee    float64   `gorm:"column:commission_fee" json:"commission_fee"`
	PlatformFee      float64   `gorm:"column:platform_fee" json:"platform_fee"`
	TransactionFee   float64   `gorm:"column:transaction_fee" json:"transaction_fee"`
	VoucherDiscount  float64   `gorm:"column:voucher_discount" json:"voucher_discount"`
	SellerIncome     float64   `gorm:"column:seller_income" json:"seller_income"`
	SubtotalInclFee  float64   `gorm:"column:subtotal_incl_fee" json:"subtotal_incl_fee"`
	TotalAmount      float64   `gorm:"column:total_amount" json:"total_amount"`
	Currency         string    `gorm:"column:currency" json:"currency"`
	Status           string    `gorm:"column:status" json:"status"`
	CreateDtm        time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateDtm        time.Time `gorm:"column:update_dtm" json:"update_dtm"`
}

func (PlatformStatementDocument) TableName() string {
	return "platform_statement_document"
}

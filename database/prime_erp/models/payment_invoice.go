package models

import (
	"time"
	"gorm.io/datatypes"
)

type PaymentInvoice struct {
	Id string `gorm:"column:id;primaryKey" json:"id"`
	PaymentId string `gorm:"column:payment_id" json:"payment_id"`
	InvoiceCode string `gorm:"column:invoice_code" json:"invoice_code"`
	Amount float64 `gorm:"column:amount" json:"amount"`
	ApplyDate time.Time `gorm:"column:apply_date" json:"apply_date"`
	CreateBy string `gorm:"column:create_by" json:"create_by"`
	CreateDtm time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateBy string `gorm:"column:update_by" json:"update_by"`
	UpdateDate time.Time `gorm:"column:update_date" json:"update_date"`
}

func (PaymentInvoice) TableName() string {
	return "payment_invoice"
}

package models

import (
	"time"
	"gorm.io/datatypes"
)

type DocumentStep struct {
	Id string `gorm:"column:id" json:"id"`
	DocumentStepCode string `gorm:"column:document_step_code" json:"document_step_code"`
	DocumentRef string `gorm:"column:document_ref" json:"document_ref"`
	PlatformCode string `gorm:"column:platform_code" json:"platform_code"`
	CurrentStep string `gorm:"column:current_step" json:"current_step"`
	CurrentStartDtm time.Time `gorm:"column:current_start_dtm" json:"current_start_dtm"`
	RefundAmount float64 `gorm:"column:refund_amount" json:"refund_amount"`
	RefundMethod string `gorm:"column:refund_method" json:"refund_method"`
	RefundDate time.Time `gorm:"column:refund_date" json:"refund_date"`
	RefundStatus string `gorm:"column:refund_status" json:"refund_status"`
	RefundCode string `gorm:"column:refund_code" json:"refund_code"`
	ShippingType string `gorm:"column:shipping_type" json:"shipping_type"`
	TrackingNo string `gorm:"column:tracking_no" json:"tracking_no"`
	ShippingStatus string `gorm:"column:shipping_status" json:"shipping_status"`
	PromisedShippingDtm time.Time `gorm:"column:promised_shipping_dtm" json:"promised_shipping_dtm"`
	PickupAddress string `gorm:"column:pickup_address" json:"pickup_address"`
	BankName string `gorm:"column:bank_name" json:"bank_name"`
	AccountName string `gorm:"column:account_name" json:"account_name"`
	BankAccountNo string `gorm:"column:bank_account_no" json:"bank_account_no"`
	BookBankUrl string `gorm:"column:book_bank_url" json:"book_bank_url"`
	CreateBy string `gorm:"column:create_by" json:"create_by"`
	CreateDtm time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateBy string `gorm:"column:update_by" json:"update_by"`
	UpdateDtm time.Time `gorm:"column:update_dtm" json:"update_dtm"`
	Shipping string `gorm:"column:shipping" json:"shipping"`
	TotalReturnPackage int32 `gorm:"column:total_return_package" json:"total_return_package"`
}

func (DocumentStep) TableName() string {
	return "document_step"
}

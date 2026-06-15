package models

import (
	"time"
	"gorm.io/datatypes"
)

type InvoiceDeposit struct {
	Id string `gorm:"column:id;primaryKey" json:"id"`
	InvoiceId string `gorm:"column:invoice_id" json:"invoice_id"`
	DepositCode string `gorm:"column:deposit_code" json:"deposit_code"`
	ApplyDate time.Time `gorm:"column:apply_date" json:"apply_date"`
	Amount float64 `gorm:"column:amount" json:"amount"`
	CreateBy string `gorm:"column:create_by" json:"create_by"`
	CreateDtm time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateBy string `gorm:"column:update_by" json:"update_by"`
	UpdateDate time.Time `gorm:"column:update_date" json:"update_date"`
}

func (InvoiceDeposit) TableName() string {
	return "invoice_deposit"
}

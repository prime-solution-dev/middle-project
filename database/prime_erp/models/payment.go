package models

import (
	"time"
)

type Payment struct {
	Id           string    `gorm:"column:id;primaryKey" json:"id"`
	PaymentCode  string    `gorm:"column:payment_code" json:"payment_code"`
	CustomerCode string    `gorm:"column:customer_code" json:"customer_code"`
	PaymentDate  time.Time `gorm:"column:payment_date" json:"payment_date"`
	Amount       float64   `gorm:"column:amount" json:"amount"`
	Method       string    `gorm:"column:method" json:"method"`
	Status       string    `gorm:"column:status" json:"status"`
	Remark       string    `gorm:"column:remark" json:"remark"`
	CreateBy     string    `gorm:"column:create_by" json:"create_by"`
	CreateDtm    time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateBy     string    `gorm:"column:update_by" json:"update_by"`
	UpdateDate   time.Time `gorm:"column:update_date" json:"update_date"`
	ExternalId   string    `gorm:"column:external_id" json:"external_id"`
}

func (Payment) TableName() string {
	return "payment"
}

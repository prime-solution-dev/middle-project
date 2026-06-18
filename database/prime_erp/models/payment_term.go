package models

import (
	"time"
)
type PaymentTerm struct {
	Id        string    `gorm:"column:id;primaryKey" json:"id"`
	TermCode  string    `gorm:"column:term_code" json:"term_code"`
	TermName  string    `gorm:"column:term_name" json:"term_name"`
	TermType  string    `gorm:"column:term_type" json:"term_type"`
	CreateBy  string    `gorm:"column:create_by" json:"create_by"`
	CreateDtm time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateBy  string    `gorm:"column:update_by" json:"update_by"`
	UpdateDtm time.Time `gorm:"column:update_dtm" json:"update_dtm"`
}

func (PaymentTerm) TableName() string {
	return "payment_term"
}

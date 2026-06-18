package models

import (
	"time"
)

type CreditExtra struct {
	Id           string    `gorm:"column:id;primaryKey" json:"id"`
	CreditId     string    `gorm:"column:credit_id" json:"credit_id"`
	ExtraType    string    `gorm:"column:extra_type" json:"extra_type"`
	Amount       float64   `gorm:"column:amount" json:"amount"`
	EffectiveDtm time.Time `gorm:"column:effective_dtm" json:"effective_dtm"`
	ExpireDtm    time.Time `gorm:"column:expire_dtm" json:"expire_dtm"`
	DocRef       string    `gorm:"column:doc_ref" json:"doc_ref"`
	ApproveDate  time.Time `gorm:"column:approve_date" json:"approve_date"`
	CreateBy     string    `gorm:"column:create_by" json:"create_by"`
	CreateDtm    time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateBy     string    `gorm:"column:update_by" json:"update_by"`
	UpdateDate   time.Time `gorm:"column:update_date" json:"update_date"`
}

func (CreditExtra) TableName() string {
	return "credit_extra"
}

package models

import (
	"time"
	"gorm.io/datatypes"
)

type Credit struct {
	Id string `gorm:"column:id;primaryKey" json:"id"`
	CustomerCode string `gorm:"column:customer_code" json:"customer_code"`
	Amount float64 `gorm:"column:amount" json:"amount"`
	EffectiveDtm time.Time `gorm:"column:effective_dtm" json:"effective_dtm"`
	IsActive bool `gorm:"column:is_active" json:"is_active"`
	DocRef string `gorm:"column:doc_ref" json:"doc_ref"`
	ApproveDate time.Time `gorm:"column:approve_date" json:"approve_date"`
	CreateBy string `gorm:"column:create_by" json:"create_by"`
	CreateDtm time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateBy string `gorm:"column:update_by" json:"update_by"`
	UpdateDate time.Time `gorm:"column:update_date" json:"update_date"`
	AlertBalanceCredit bool `gorm:"column:alert_balance_credit" json:"alert_balance_credit"`
}

func (Credit) TableName() string {
	return "credit"
}

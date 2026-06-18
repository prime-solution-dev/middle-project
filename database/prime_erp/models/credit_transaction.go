package models

import (
	"time"
)

type CreditTransaction struct {
	Id              string    `gorm:"column:id;primaryKey" json:"id"`
	TransactionCode string    `gorm:"column:transaction_code" json:"transaction_code"`
	TransactionType string    `gorm:"column:transaction_type" json:"transaction_type"`
	Amount          float64   `gorm:"column:amount" json:"amount"`
	AdjustAmount    float64   `gorm:"column:adjust_amount" json:"adjust_amount"`
	EffectiveDtm    time.Time `gorm:"column:effective_dtm" json:"effective_dtm"`
	ExpireDtm       time.Time `gorm:"column:expire_dtm" json:"expire_dtm"`
	ForceExpireDtm  time.Time `gorm:"column:force_expire_dtm" json:"force_expire_dtm"`
	IsApprove       bool      `gorm:"column:is_approve" json:"is_approve"`
	Status          string    `gorm:"column:status" json:"status"`
	Reason          string    `gorm:"column:reason" json:"reason"`
	ApproveDate     time.Time `gorm:"column:approve_date" json:"approve_date"`
	CreateBy        string    `gorm:"column:create_by" json:"create_by"`
	CreateDtm       time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateBy        string    `gorm:"column:update_by" json:"update_by"`
	UpdateDate      time.Time `gorm:"column:update_date" json:"update_date"`
}

func (CreditTransaction) TableName() string {
	return "credit_transaction"
}

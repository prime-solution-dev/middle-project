package models

import (
	"time"
)

type Deposit struct {
	Id           string    `gorm:"column:id;primaryKey" json:"id"`
	DepositCode  string    `gorm:"column:deposit_code" json:"deposit_code"`
	DocRefType   string    `gorm:"column:doc_ref_type" json:"doc_ref_type"`
	DocRef       string    `gorm:"column:doc_ref" json:"doc_ref"`
	CustomerCode string    `gorm:"column:customer_code" json:"customer_code"`
	DepositDate  time.Time `gorm:"column:deposit_date" json:"deposit_date"`
	AmountTotal  float64   `gorm:"column:amount_total" json:"amount_total"`
	AmountUsed   float64   `gorm:"column:amount_used" json:"amount_used"`
	Status       string    `gorm:"column:status" json:"status"`
	Remark       string    `gorm:"column:remark" json:"remark"`
	CreateBy     string    `gorm:"column:create_by" json:"create_by"`
	CreateDtm    time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateBy     string    `gorm:"column:update_by" json:"update_by"`
	UpdateDate   time.Time `gorm:"column:update_date" json:"update_date"`
	AmountRemain float64   `gorm:"column:amount_remain" json:"amount_remain"`
}

func (Deposit) TableName() string {
	return "deposit"
}

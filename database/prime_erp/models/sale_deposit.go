package models

import (
	"time"
	"gorm.io/datatypes"
)

type SaleDeposit struct {
	Id string `gorm:"column:id;primaryKey" json:"id"`
	SaleId string `gorm:"column:sale_id" json:"sale_id"`
	DepositCode string `gorm:"column:deposit_code" json:"deposit_code"`
	DepositDate time.Time `gorm:"column:deposit_date" json:"deposit_date"`
	AmountTotal float64 `gorm:"column:amount_total" json:"amount_total"`
	AmountUsed float64 `gorm:"column:amount_used" json:"amount_used"`
	AmountRemain float64 `gorm:"column:amount_remain" json:"amount_remain"`
}

func (SaleDeposit) TableName() string {
	return "sale_deposit"
}

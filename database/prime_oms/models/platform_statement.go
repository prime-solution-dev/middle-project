package models

import (
	"time"
	"gorm.io/datatypes"
)

type PlatformStatement struct {
	Id string `gorm:"column:id" json:"id"`
	ShopCode string `gorm:"column:shop_code" json:"shop_code"`
	StatementCode string `gorm:"column:statement_code" json:"statement_code"`
	PlatformCode string `gorm:"column:platform_code" json:"platform_code"`
	StatementType string `gorm:"column:statement_type" json:"statement_type"`
	DocumentRef string `gorm:"column:document_ref" json:"document_ref"`
	PeriodFromDtm time.Time `gorm:"column:period_from_dtm" json:"period_from_dtm"`
	PeriodToDtm time.Time `gorm:"column:period_to_dtm" json:"period_to_dtm"`
	TotalAmount float64 `gorm:"column:total_amount" json:"total_amount"`
	TotalFeeAmount float64 `gorm:"column:total_fee_amount" json:"total_fee_amount"`
	Currency string `gorm:"column:currency" json:"currency"`
	Status string `gorm:"column:status" json:"status"`
	CreateDtm time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateDtm time.Time `gorm:"column:update_dtm" json:"update_dtm"`
}

func (PlatformStatement) TableName() string {
	return "platform_statement"
}

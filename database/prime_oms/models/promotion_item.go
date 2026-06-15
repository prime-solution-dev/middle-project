package models

import (
	"time"
	"gorm.io/datatypes"
)

type PromotionItem struct {
	Id string `gorm:"column:id" json:"id"`
	PromoId string `gorm:"column:promo_id" json:"promo_id"`
	PromoItem string `gorm:"column:promo_item" json:"promo_item"`
	Qty float64 `gorm:"column:qty" json:"qty"`
	QtyBenefit float64 `gorm:"column:qty_benefit" json:"qty_benefit"`
	QtyLimit int64 `gorm:"column:qty_limit" json:"qty_limit"`
	MinAmount float64 `gorm:"column:min_amount" json:"min_amount"`
	MaxAmount float64 `gorm:"column:max_amount" json:"max_amount"`
	Seq int64 `gorm:"column:seq" json:"seq"`
	ConditionType string `gorm:"column:condition_type" json:"condition_type"`
	Status string `gorm:"column:status" json:"status"`
	CreateDtm time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateDtm time.Time `gorm:"column:update_dtm" json:"update_dtm"`
	CreateBy string `gorm:"column:create_by" json:"create_by"`
	UpdateBy string `gorm:"column:update_by" json:"update_by"`
}

func (PromotionItem) TableName() string {
	return "promotion_item"
}

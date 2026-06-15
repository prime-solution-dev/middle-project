package models

import (
	"time"
	"gorm.io/datatypes"
)

type PromotionItemUsage struct {
	Id string `gorm:"column:id" json:"id"`
	PromotionCode string `gorm:"column:promotion_code" json:"promotion_code"`
	PromotionItem string `gorm:"column:promotion_item" json:"promotion_item"`
	PromoConditionItem string `gorm:"column:promo_condition_item" json:"promo_condition_item"`
	Status string `gorm:"column:status" json:"status"`
	QtyPromo float64 `gorm:"column:qty_promo" json:"qty_promo"`
	QtyBenefit float64 `gorm:"column:qty_benefit" json:"qty_benefit"`
	SaleCode string `gorm:"column:sale_code" json:"sale_code"`
	PromoType string `gorm:"column:promo_type" json:"promo_type"`
	SnapshotSource datatypes.JSON `gorm:"column:snapshot_source" json:"snapshot_source"`
	SnapshotDestination datatypes.JSON `gorm:"column:snapshot_destination" json:"snapshot_destination"`
	CreateBy string `gorm:"column:create_by" json:"create_by"`
	CreateDtm time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateDtm time.Time `gorm:"column:update_dtm" json:"update_dtm"`
}

func (PromotionItemUsage) TableName() string {
	return "promotion_item_usage"
}

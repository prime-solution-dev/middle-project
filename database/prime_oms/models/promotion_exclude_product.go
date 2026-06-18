package models

import (
	"time"
)

type PromotionExcludeProduct struct {
	Id          string    `gorm:"column:id" json:"id"`
	PromoId     string    `gorm:"column:promo_id" json:"promo_id"`
	ProductCode string    `gorm:"column:product_code" json:"product_code"`
	CreateDtm   time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateDtm   time.Time `gorm:"column:update_dtm" json:"update_dtm"`
	CreateBy    string    `gorm:"column:create_by" json:"create_by"`
	UpdateBy    string    `gorm:"column:update_by" json:"update_by"`
}

func (PromotionExcludeProduct) TableName() string {
	return "promotion_exclude_product"
}

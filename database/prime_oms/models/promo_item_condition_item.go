package models

import (
	"time"
	"gorm.io/datatypes"
)

type PromoItemConditionItem struct {
	Id string `gorm:"column:id" json:"id"`
	PromoItemConditionId string `gorm:"column:promo_item_condition_id" json:"promo_item_condition_id"`
	ProductCode string `gorm:"column:product_code" json:"product_code"`
	Qty int32 `gorm:"column:qty" json:"qty"`
	QtyLimit int32 `gorm:"column:qty_limit" json:"qty_limit"`
	Status string `gorm:"column:status" json:"status"`
	Seq int32 `gorm:"column:seq" json:"seq"`
	CreateDtm time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateDtm time.Time `gorm:"column:update_dtm" json:"update_dtm"`
}

func (PromoItemConditionItem) TableName() string {
	return "promo_item_condition_item"
}

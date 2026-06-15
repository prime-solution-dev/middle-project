package models

import (
	"time"
	"gorm.io/datatypes"
)

type PromoItemCondition struct {
	Id string `gorm:"column:id" json:"id"`
	PromoItemId string `gorm:"column:promo_item_id" json:"promo_item_id"`
	ConditionType string `gorm:"column:condition_type" json:"condition_type"`
	Seq int32 `gorm:"column:seq" json:"seq"`
}

func (PromoItemCondition) TableName() string {
	return "promo_item_condition"
}

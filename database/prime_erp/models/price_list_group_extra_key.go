package models

import (
	"time"
	"gorm.io/datatypes"
)

type PriceListGroupExtraKey struct {
	Id string `gorm:"column:id;primaryKey" json:"id"`
	GroupExtraId string `gorm:"column:group_extra_id" json:"group_extra_id"`
	Code string `gorm:"column:code" json:"code"`
	Value string `gorm:"column:value" json:"value"`
	Seq int32 `gorm:"column:seq" json:"seq"`
}

func (PriceListGroupExtraKey) TableName() string {
	return "price_list_group_extra_key"
}

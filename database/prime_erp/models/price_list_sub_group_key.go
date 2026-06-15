package models

import (
	"time"
	"gorm.io/datatypes"
)

type PriceListSubGroupKey struct {
	Id string `gorm:"column:id;primaryKey" json:"id"`
	SubGroupId string `gorm:"column:sub_group_id" json:"sub_group_id"`
	Code string `gorm:"column:code" json:"code"`
	Value string `gorm:"column:value" json:"value"`
	Seq int32 `gorm:"column:seq" json:"seq"`
}

func (PriceListSubGroupKey) TableName() string {
	return "price_list_sub_group_key"
}

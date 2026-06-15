package models

import (
	"time"
	"gorm.io/datatypes"
)

type PriceListSubGroupKeyHistory struct {
	Id string `gorm:"column:id;primaryKey" json:"id"`
	SubGroupHistoryId string `gorm:"column:sub_group_history_id" json:"sub_group_history_id"`
	Code string `gorm:"column:code" json:"code"`
	Value string `gorm:"column:value" json:"value"`
	Seq int32 `gorm:"column:seq" json:"seq"`
}

func (PriceListSubGroupKeyHistory) TableName() string {
	return "price_list_sub_group_key_history"
}

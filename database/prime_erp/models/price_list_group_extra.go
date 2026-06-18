package models

import (
	"time"
)

type PriceListGroupExtra struct {
	Id               string    `gorm:"column:id;primaryKey" json:"id"`
	PriceListGroupId string    `gorm:"column:price_list_group_id" json:"price_list_group_id"`
	ExtraKey         string    `gorm:"column:extra_key" json:"extra_key"`
	LengthExtraKey   int32     `gorm:"column:length_extra_key" json:"length_extra_key"`
	Operator         string    `gorm:"column:operator" json:"operator"`
	CondRangeMin     float64   `gorm:"column:cond_range_min" json:"cond_range_min"`
	CondRangeMax     float64   `gorm:"column:cond_range_max" json:"cond_range_max"`
	CreateBy         string    `gorm:"column:create_by" json:"create_by"`
	CreateDtm        time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateBy         string    `gorm:"column:update_by" json:"update_by"`
	UpdateDtm        time.Time `gorm:"column:update_dtm" json:"update_dtm"`
	ConditionCode    string    `gorm:"column:condition_code" json:"condition_code"`
	ValueInt         int32     `gorm:"column:value_int" json:"value_int"`
}

func (PriceListGroupExtra) TableName() string {
	return "price_list_group_extra"
}

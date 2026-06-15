package models

import (
	"time"
	"gorm.io/datatypes"
)

type PriceListGroupTerm struct {
	Id string `gorm:"column:id;primaryKey" json:"id"`
	PriceListGroupId string `gorm:"column:price_list_group_id" json:"price_list_group_id"`
	TermCode string `gorm:"column:term_code" json:"term_code"`
	Pdc float64 `gorm:"column:pdc" json:"pdc"`
	PdcPercent int32 `gorm:"column:pdc_percent" json:"pdc_percent"`
	Due float64 `gorm:"column:due" json:"due"`
	DuePercent int32 `gorm:"column:due_percent" json:"due_percent"`
	CreateBy string `gorm:"column:create_by" json:"create_by"`
	CreateDtm time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateBy string `gorm:"column:update_by" json:"update_by"`
	UpdateDtm time.Time `gorm:"column:update_dtm" json:"update_dtm"`
}

func (PriceListGroupTerm) TableName() string {
	return "price_list_group_term"
}

package models

import (
	"time"
	"gorm.io/datatypes"
)

type PriceListGroup struct {
	Id string `gorm:"column:id;primaryKey" json:"id"`
	CompanyCode string `gorm:"column:company_code" json:"company_code"`
	SiteCode string `gorm:"column:site_code" json:"site_code"`
	GroupCode string `gorm:"column:group_code" json:"group_code"`
	PriceUnit float64 `gorm:"column:price_unit" json:"price_unit"`
	PriceWeight float64 `gorm:"column:price_weight" json:"price_weight"`
	BeforePriceUnit float64 `gorm:"column:before_price_unit" json:"before_price_unit"`
	BeforePriceWeight float64 `gorm:"column:before_price_weight" json:"before_price_weight"`
	Currency string `gorm:"column:currency" json:"currency"`
	EffectiveDate time.Time `gorm:"column:effective_date" json:"effective_date"`
	ExtraPattern string `gorm:"column:extra_pattern" json:"extra_pattern"`
	Remark string `gorm:"column:remark" json:"remark"`
	CreateBy string `gorm:"column:create_by" json:"create_by"`
	CreateDtm time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateBy string `gorm:"column:update_by" json:"update_by"`
	UpdateDtm time.Time `gorm:"column:update_dtm" json:"update_dtm"`
	GroupKey string `gorm:"column:group_key" json:"group_key"`
	GroupName string `gorm:"column:group_name" json:"group_name"`
}

func (PriceListGroup) TableName() string {
	return "price_list_group"
}

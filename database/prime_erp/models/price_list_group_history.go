package models

import (
	"time"
)

type PriceListGroupHistory struct {
	Id                string    `gorm:"column:id;primaryKey" json:"id"`
	CompanyCode       string    `gorm:"column:company_code" json:"company_code"`
	SiteCode          string    `gorm:"column:site_code" json:"site_code"`
	GroupCode         string    `gorm:"column:group_code" json:"group_code"`
	PriceUnit         float64   `gorm:"column:price_unit" json:"price_unit"`
	PriceWeight       float64   `gorm:"column:price_weight" json:"price_weight"`
	BeforePriceUnit   float64   `gorm:"column:before_price_unit" json:"before_price_unit"`
	BeforePriceWeight float64   `gorm:"column:before_price_weight" json:"before_price_weight"`
	Currency          string    `gorm:"column:currency" json:"currency"`
	EffectiveDate     time.Time `gorm:"column:effective_date" json:"effective_date"`
	ExpiryDate        time.Time `gorm:"column:expiry_date" json:"expiry_date"`
	ExtraPattern      string    `gorm:"column:extra_pattern" json:"extra_pattern"`
	Remark            string    `gorm:"column:remark" json:"remark"`
	CreateBy          string    `gorm:"column:create_by" json:"create_by"`
	CreateDtm         time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateBy          string    `gorm:"column:update_by" json:"update_by"`
	UpdateDtm         time.Time `gorm:"column:update_dtm" json:"update_dtm"`
	PriceListGroupId  string    `gorm:"column:price_list_group_id" json:"price_list_group_id"`
}

func (PriceListGroupHistory) TableName() string {
	return "price_list_group_history"
}

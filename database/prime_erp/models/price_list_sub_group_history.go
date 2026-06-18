package models

import (
	"time"
)

type PriceListSubGroupHistory struct {
	Id                        string    `gorm:"column:id;primaryKey" json:"id"`
	PriceListGroupId          string    `gorm:"column:price_list_group_id" json:"price_list_group_id"`
	SubgroupKey               string    `gorm:"column:subgroup_key" json:"subgroup_key"`
	IsTrading                 bool      `gorm:"column:is_trading" json:"is_trading"`
	PriceUnit                 float64   `gorm:"column:price_unit" json:"price_unit"`
	ExtraPriceUnit            float64   `gorm:"column:extra_price_unit" json:"extra_price_unit"`
	TotalNetPriceUnit         float64   `gorm:"column:total_net_price_unit" json:"total_net_price_unit"`
	PriceWeight               float64   `gorm:"column:price_weight" json:"price_weight"`
	ExtraPriceWeight          float64   `gorm:"column:extra_price_weight" json:"extra_price_weight"`
	TermPriceWeight           float64   `gorm:"column:term_price_weight" json:"term_price_weight"`
	TotalNetPriceWeight       float64   `gorm:"column:total_net_price_weight" json:"total_net_price_weight"`
	BeforePriceUnit           float64   `gorm:"column:before_price_unit" json:"before_price_unit"`
	BeforeExtraPriceUnit      float64   `gorm:"column:before_extra_price_unit" json:"before_extra_price_unit"`
	BeforeTermPriceUnit       float64   `gorm:"column:before_term_price_unit" json:"before_term_price_unit"`
	BeforeTotalNetPriceUnit   float64   `gorm:"column:before_total_net_price_unit" json:"before_total_net_price_unit"`
	BeforePriceWeight         float64   `gorm:"column:before_price_weight" json:"before_price_weight"`
	BeforeExtraPriceWeight    float64   `gorm:"column:before_extra_price_weight" json:"before_extra_price_weight"`
	BeforeTermPriceWeight     float64   `gorm:"column:before_term_price_weight" json:"before_term_price_weight"`
	BeforeTotalNetPriceWeight float64   `gorm:"column:before_total_net_price_weight" json:"before_total_net_price_weight"`
	EffectiveDate             time.Time `gorm:"column:effective_date" json:"effective_date"`
	ExpiryDate                time.Time `gorm:"column:expiry_date" json:"expiry_date"`
	Remark                    string    `gorm:"column:remark" json:"remark"`
	CreateBy                  string    `gorm:"column:create_by" json:"create_by"`
	CreateDtm                 time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateBy                  string    `gorm:"column:update_by" json:"update_by"`
	UpdateDtm                 time.Time `gorm:"column:update_dtm" json:"update_dtm"`
}

func (PriceListSubGroupHistory) TableName() string {
	return "price_list_sub_group_history"
}

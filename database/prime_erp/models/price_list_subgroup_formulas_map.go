package models

import (
	"time"
)

type PriceListSubgroupFormulasMap struct {
	Id                    string    `gorm:"column:id;primaryKey" json:"id"`
	PriceListSubgroupCode string    `gorm:"column:price_list_subgroup_code" json:"price_list_subgroup_code"`
	PriceListFormulasCode string    `gorm:"column:price_list_formulas_code" json:"price_list_formulas_code"`
	IsDefault             bool      `gorm:"column:is_default" json:"is_default"`
	CreateDtm             time.Time `gorm:"column:create_dtm" json:"create_dtm"`
}

func (PriceListSubgroupFormulasMap) TableName() string {
	return "price_list_subgroup_formulas_map"
}

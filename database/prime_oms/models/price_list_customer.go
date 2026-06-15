package models

import (
	"time"
	"gorm.io/datatypes"
)

type PriceListCustomer struct {
	Id string `gorm:"column:id" json:"id"`
	ProductCode string `gorm:"column:product_code" json:"product_code"`
	CustomerCode string `gorm:"column:customer_code" json:"customer_code"`
	ExternalProductCode string `gorm:"column:external_product_code" json:"external_product_code"`
	Currency string `gorm:"column:currency" json:"currency"`
	EffectiveDate time.Time `gorm:"column:effective_date" json:"effective_date"`
	ExpireDate time.Time `gorm:"column:expire_date" json:"expire_date"`
	ConverterQty float64 `gorm:"column:converter_qty" json:"converter_qty"`
	SpecialFactorQty float64 `gorm:"column:special_factor_qty" json:"special_factor_qty"`
	PriceMethod string `gorm:"column:price_method" json:"price_method"`
	PriceMethodValue string `gorm:"column:price_method_value" json:"price_method_value"`
	PriceUnit float64 `gorm:"column:price_unit" json:"price_unit"`
	TotalPriceAdjustment float64 `gorm:"column:total_price_adjustment" json:"total_price_adjustment"`
	TotalVat float64 `gorm:"column:total_vat" json:"total_vat"`
	TotalNetPriceUnit float64 `gorm:"column:total_net_price_unit" json:"total_net_price_unit"`
	VatType string `gorm:"column:vat_type" json:"vat_type"`
	SoldToCode string `gorm:"column:sold_to_code" json:"sold_to_code"`
	SoldTo string `gorm:"column:sold_to" json:"sold_to"`
	ShipToCode string `gorm:"column:ship_to_code" json:"ship_to_code"`
	ShipTo string `gorm:"column:ship_to" json:"ship_to"`
	Remark string `gorm:"column:remark" json:"remark"`
	IsActive bool `gorm:"column:is_active" json:"is_active"`
	CreateBy string `gorm:"column:create_by" json:"create_by"`
	CreateDtm time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateBy string `gorm:"column:update_by" json:"update_by"`
	UpdateDtm time.Time `gorm:"column:update_dtm" json:"update_dtm"`
}

func (PriceListCustomer) TableName() string {
	return "price_list_customer"
}

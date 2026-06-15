package models

import (
	"time"
	"gorm.io/datatypes"
)

type SaleItem struct {
	Id string `gorm:"column:id;primaryKey" json:"id"`
	SaleId string `gorm:"column:sale_id" json:"sale_id"`
	SaleItem string `gorm:"column:sale_item" json:"sale_item"`
	ProductCode string `gorm:"column:product_code" json:"product_code"`
	ProductDesc string `gorm:"column:product_desc" json:"product_desc"`
	Qty float64 `gorm:"column:qty" json:"qty"`
	OriginQty float64 `gorm:"column:origin_qty" json:"origin_qty"`
	Unit string `gorm:"column:unit" json:"unit"`
	PriceListUnit float64 `gorm:"column:price_list_unit" json:"price_list_unit"`
	SaleQty float64 `gorm:"column:sale_qty" json:"sale_qty"`
	SaleUnit string `gorm:"column:sale_unit" json:"sale_unit"`
	SaleUnitType string `gorm:"column:sale_unit_type" json:"sale_unit_type"`
	PassPriceUnit string `gorm:"column:pass_price_unit" json:"pass_price_unit"`
	PassPriceWeight string `gorm:"column:pass_price_weight" json:"pass_price_weight"`
	PriceUnit float64 `gorm:"column:price_unit" json:"price_unit"`
	TotalAmount float64 `gorm:"column:total_amount" json:"total_amount"`
	TransportCostUnit float64 `gorm:"column:transport_cost_unit" json:"transport_cost_unit"`
	SubtotalExclTransport float64 `gorm:"column:subtotal_excl_transport" json:"subtotal_excl_transport"`
	NetPriceUnitExclTransport float64 `gorm:"column:net_price_unit_excl_transport" json:"net_price_unit_excl_transport"`
	WeightUnit float64 `gorm:"column:weight_unit" json:"weight_unit"`
	AvgWeightUnit float64 `gorm:"column:avg_weight_unit" json:"avg_weight_unit"`
	TotalWeight float64 `gorm:"column:total_weight" json:"total_weight"`
	TransportCostWeightUnit float64 `gorm:"column:transport_cost_weight_unit" json:"transport_cost_weight_unit"`
	SubtotalWeightExclTransport float64 `gorm:"column:subtotal_weight_excl_transport" json:"subtotal_weight_excl_transport"`
	NetPricePerWeightExclTransport float64 `gorm:"column:net_price_per_weight_excl_transport" json:"net_price_per_weight_excl_transport"`
	Status string `gorm:"column:status" json:"status"`
	Remark string `gorm:"column:remark" json:"remark"`
	CreateDate time.Time `gorm:"column:create_date" json:"create_date"`
	CreateBy string `gorm:"column:create_by" json:"create_by"`
	UpdateDate time.Time `gorm:"column:update_date" json:"update_date"`
	UpdateBy string `gorm:"column:update_by" json:"update_by"`
	SubtotalExclVat float64 `gorm:"column:subtotal_excl_vat" json:"subtotal_excl_vat"`
	TotalVat float64 `gorm:"column:total_vat" json:"total_vat"`
	UnitUom string `gorm:"column:unit_uom" json:"unit_uom"`
	TotalDiscount float64 `gorm:"column:total_discount" json:"total_discount"`
	TotalDiscountPercent float64 `gorm:"column:total_discount_percent" json:"total_discount_percent"`
	DocumentRef string `gorm:"column:document_ref" json:"document_ref"`
	DocumentRefItem string `gorm:"column:document_ref_item" json:"document_ref_item"`
	OldPriceListUnit float64 `gorm:"column:old_price_list_unit" json:"old_price_list_unit"`
}

func (SaleItem) TableName() string {
	return "sale_item"
}

package models

import (
	"time"
)

type SaleItem struct {
	Id                             string    `gorm:"column:id" json:"id"`
	SaleId                         string    `gorm:"column:sale_id" json:"sale_id"`
	SaleItem                       string    `gorm:"column:sale_item" json:"sale_item"`
	ProductCode                    string    `gorm:"column:product_code" json:"product_code"`
	ProductDesc                    string    `gorm:"column:product_desc" json:"product_desc"`
	Qty                            float64   `gorm:"column:qty" json:"qty"`
	OriginQty                      float64   `gorm:"column:origin_qty" json:"origin_qty"`
	Unit                           string    `gorm:"column:unit" json:"unit"`
	PriceListUnit                  float64   `gorm:"column:price_list_unit" json:"price_list_unit"`
	SaleQty                        float64   `gorm:"column:sale_qty" json:"sale_qty"`
	SaleUnit                       string    `gorm:"column:sale_unit" json:"sale_unit"`
	SaleUnitType                   string    `gorm:"column:sale_unit_type" json:"sale_unit_type"`
	PassPriceUnit                  string    `gorm:"column:pass_price_unit" json:"pass_price_unit"`
	PassPriceWeight                string    `gorm:"column:pass_price_weight" json:"pass_price_weight"`
	PriceUnit                      float64   `gorm:"column:price_unit" json:"price_unit"`
	TotalAmount                    float64   `gorm:"column:total_amount" json:"total_amount"`
	SubtotalExclVat                float64   `gorm:"column:subtotal_excl_vat" json:"subtotal_excl_vat"`
	TotalVat                       float64   `gorm:"column:total_vat" json:"total_vat"`
	TransportCostUnit              float64   `gorm:"column:transport_cost_unit" json:"transport_cost_unit"`
	SubtotalExclTransport          float64   `gorm:"column:subtotal_excl_transport" json:"subtotal_excl_transport"`
	NetPriceUnitExclTransport      float64   `gorm:"column:net_price_unit_excl_transport" json:"net_price_unit_excl_transport"`
	WeightUnit                     float64   `gorm:"column:weight_unit" json:"weight_unit"`
	AvgWeightUnit                  float64   `gorm:"column:avg_weight_unit" json:"avg_weight_unit"`
	TotalWeight                    float64   `gorm:"column:total_weight" json:"total_weight"`
	TransportCostWeightUnit        float64   `gorm:"column:transport_cost_weight_unit" json:"transport_cost_weight_unit"`
	SubtotalWeightExclTransport    float64   `gorm:"column:subtotal_weight_excl_transport" json:"subtotal_weight_excl_transport"`
	NetPricePerWeightExclTransport float64   `gorm:"column:net_price_per_weight_excl_transport" json:"net_price_per_weight_excl_transport"`
	Status                         string    `gorm:"column:status" json:"status"`
	Remark                         string    `gorm:"column:remark" json:"remark"`
	UnitUom                        string    `gorm:"column:unit_uom" json:"unit_uom"`
	TotalDiscount                  float64   `gorm:"column:total_discount" json:"total_discount"`
	TotalDiscountPercent           float64   `gorm:"column:total_discount_percent" json:"total_discount_percent"`
	CreateDate                     time.Time `gorm:"column:create_date" json:"create_date"`
	CreateBy                       string    `gorm:"column:create_by" json:"create_by"`
	UpdateDate                     time.Time `gorm:"column:update_date" json:"update_date"`
	UpdateBy                       string    `gorm:"column:update_by" json:"update_by"`
	PromotionId                    string    `gorm:"column:promotion_id" json:"promotion_id"`
	ItemType                       string    `gorm:"column:item_type" json:"item_type"`
	ParentItemId                   string    `gorm:"column:parent_item_id" json:"parent_item_id"`
	PromotionUsageId               string    `gorm:"column:promotion_usage_id" json:"promotion_usage_id"`
	ActualQty                      float64   `gorm:"column:actual_qty" json:"actual_qty"`
	PlanQty                        float64   `gorm:"column:plan_qty" json:"plan_qty"`
	DocumentRefItem                string    `gorm:"column:document_ref_item" json:"document_ref_item"`
	InterfaceQty                   float64   `gorm:"column:interface_qty" json:"interface_qty"`
	SaleBomParentCode              string    `gorm:"column:sale_bom_parent_code" json:"sale_bom_parent_code"`
	DiscountUnitVat                float64   `gorm:"column:discount_unit_vat" json:"discount_unit_vat"`
	TotalDiscountVat               float64   `gorm:"column:total_discount_vat" json:"total_discount_vat"`
	SubtotalAmountVat              float64   `gorm:"column:subtotal_amount_vat" json:"subtotal_amount_vat"`
	SubtotalAmount                 float64   `gorm:"column:subtotal_amount" json:"subtotal_amount"`
	TotalAmound                    float64   `gorm:"column:total_amound" json:"total_amound"`
	PriceUnitVat                   float64   `gorm:"column:price_unit_vat" json:"price_unit_vat"`
}

func (SaleItem) TableName() string {
	return "sale_item"
}

package models

import (
	"time"
)

type PrePurchaseItem struct {
	Id                   string    `gorm:"column:id;primaryKey" json:"id"`
	PrePurchaseId        string    `gorm:"column:pre_purchase_id" json:"pre_purchase_id"`
	PreItem              string    `gorm:"column:pre_item" json:"pre_item"`
	HierarchyType        string    `gorm:"column:hierarchy_type" json:"hierarchy_type"`
	HierarchyCode        string    `gorm:"column:hierarchy_code" json:"hierarchy_code"`
	DocRefItem           string    `gorm:"column:doc_ref_item" json:"doc_ref_item"`
	Qty                  float64   `gorm:"column:qty" json:"qty"`
	Unit                 string    `gorm:"column:unit" json:"unit"`
	PurchaseQty          float64   `gorm:"column:purchase_qty" json:"purchase_qty"`
	PurchaseUnit         string    `gorm:"column:purchase_unit" json:"purchase_unit"`
	PurchaseUnitType     string    `gorm:"column:purchase_unit_type" json:"purchase_unit_type"`
	PriceUnit            float64   `gorm:"column:price_unit" json:"price_unit"`
	TotalDiscount        float64   `gorm:"column:total_discount" json:"total_discount"`
	TotalAmount          float64   `gorm:"column:total_amount" json:"total_amount"`
	UnitUom              string    `gorm:"column:unit_uom" json:"unit_uom"`
	TotalCost            float64   `gorm:"column:total_cost" json:"total_cost"`
	TotalDiscountPercent float64   `gorm:"column:total_discount_percent" json:"total_discount_percent"`
	TotalVat             float64   `gorm:"column:total_vat" json:"total_vat"`
	SubtotalExclVat      float64   `gorm:"column:subtotal_excl_vat" json:"subtotal_excl_vat"`
	WeightUnit           float64   `gorm:"column:weight_unit" json:"weight_unit"`
	TotalWeight          float64   `gorm:"column:total_weight" json:"total_weight"`
	Status               string    `gorm:"column:status" json:"status"`
	Remark               string    `gorm:"column:remark" json:"remark"`
	CreateDtm            time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	CreateBy             string    `gorm:"column:create_by" json:"create_by"`
	UpdateDtm            time.Time `gorm:"column:update_dtm" json:"update_dtm"`
	UpdateBy             string    `gorm:"column:update_by" json:"update_by"`
	DiscountType         string    `gorm:"column:discount_type" json:"discount_type"`
}

func (PrePurchaseItem) TableName() string {
	return "pre_purchase_item"
}

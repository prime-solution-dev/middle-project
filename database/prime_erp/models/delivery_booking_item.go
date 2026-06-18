package models

import (
	"time"
)

type DeliveryBookingItem struct {
	Id              string    `gorm:"column:id;primaryKey" json:"id"`
	DeliveryItem    string    `gorm:"column:delivery_item" json:"delivery_item"`
	DeliveryId      string    `gorm:"column:delivery_id" json:"delivery_id"`
	ProductCode     string    `gorm:"column:product_code" json:"product_code"`
	Qty             float64   `gorm:"column:qty" json:"qty"`
	UnitCode        string    `gorm:"column:unit_code" json:"unit_code"`
	PriceListUnit   float64   `gorm:"column:price_list_unit" json:"price_list_unit"`
	SaleQty         float64   `gorm:"column:sale_qty" json:"sale_qty"`
	SaleUnitCode    string    `gorm:"column:sale_unit_code" json:"sale_unit_code"`
	TotalWeight     float64   `gorm:"column:total_weight" json:"total_weight"`
	CreateDate      time.Time `gorm:"column:create_date" json:"create_date"`
	CreateBy        string    `gorm:"column:create_by" json:"create_by"`
	UpdateDate      time.Time `gorm:"column:update_date" json:"update_date"`
	UpdateBy        string    `gorm:"column:update_by" json:"update_by"`
	DocumentRefItem string    `gorm:"column:document_ref_item" json:"document_ref_item"`
	Status          string    `gorm:"column:status" json:"status"`
	Weight          float64   `gorm:"column:weight" json:"weight"`
	WeightUnit      float64   `gorm:"column:weight_unit" json:"weight_unit"`
	Remark          string    `gorm:"column:remark" json:"remark"`
}

func (DeliveryBookingItem) TableName() string {
	return "delivery_booking_item"
}

package models

import (
	"time"
)

type ShipmentItem struct {
	Id              string    `gorm:"column:id" json:"id"`
	ShipmentId      string    `gorm:"column:shipment_id" json:"shipment_id"`
	ShipmentItem    string    `gorm:"column:shipment_item" json:"shipment_item"`
	DocumentRef     string    `gorm:"column:document_ref" json:"document_ref"`
	DocumentItemRef string    `gorm:"column:document_item_ref" json:"document_item_ref"`
	SaleCode        string    `gorm:"column:sale_code" json:"sale_code"`
	SaleItem        string    `gorm:"column:sale_item" json:"sale_item"`
	ProductCode     string    `gorm:"column:product_code" json:"product_code"`
	Qty             float64   `gorm:"column:qty" json:"qty"`
	Status          string    `gorm:"column:status" json:"status"`
	AssignedBy      string    `gorm:"column:assigned_by" json:"assigned_by"`
	CreateBy        string    `gorm:"column:create_by" json:"create_by"`
	CreateDtm       time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateBy        string    `gorm:"column:update_by" json:"update_by"`
	UpdateDtm       time.Time `gorm:"column:update_dtm" json:"update_dtm"`
}

func (ShipmentItem) TableName() string {
	return "shipment_item"
}

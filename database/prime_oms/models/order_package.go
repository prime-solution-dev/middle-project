package models

import (
	"time"
)

type OrderPackage struct {
	Id             string    `gorm:"column:id" json:"id"`
	PackageCode    string    `gorm:"column:package_code" json:"package_code"`
	DocumentRef    string    `gorm:"column:document_ref" json:"document_ref"`
	CourierCode    string    `gorm:"column:courier_code" json:"courier_code"`
	TrackingNo     string    `gorm:"column:tracking_no" json:"tracking_no"`
	Status         string    `gorm:"column:status" json:"status"`
	StatusShipment string    `gorm:"column:status_shipment" json:"status_shipment"`
	Weight         float64   `gorm:"column:weight" json:"weight"`
	Dimension      float64   `gorm:"column:dimension" json:"dimension"`
	DimensionDesc  string    `gorm:"column:dimension_desc" json:"dimension_desc"`
	LabelUrl       string    `gorm:"column:label_url" json:"label_url"`
	ShipDate       time.Time `gorm:"column:ship_date" json:"ship_date"`
	CreateBy       string    `gorm:"column:create_by" json:"create_by"`
	CreateDtm      time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateBy       string    `gorm:"column:update_by" json:"update_by"`
	UpdateDtm      time.Time `gorm:"column:update_dtm" json:"update_dtm"`
	ExternalRef    string    `gorm:"column:external_ref" json:"external_ref"`
	ShopCode       string    `gorm:"column:shop_code" json:"shop_code"`
	IsCod          bool      `gorm:"column:is_cod" json:"is_cod"`
}

func (OrderPackage) TableName() string {
	return "order_package"
}

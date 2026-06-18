package models

import (
	"time"
)

type Shipment struct {
	Id                 string    `gorm:"column:id" json:"id"`
	ShipmentCode       string    `gorm:"column:shipment_code" json:"shipment_code"`
	CourierCode        string    `gorm:"column:courier_code" json:"courier_code"`
	ShipmentType       string    `gorm:"column:shipment_type" json:"shipment_type"`
	TrackingNo         string    `gorm:"column:tracking_no" json:"tracking_no"`
	Status             string    `gorm:"column:status" json:"status"`
	DeliveryOptionId   string    `gorm:"column:delivery_option_id" json:"delivery_option_id"`
	ShippingProviderId string    `gorm:"column:shipping_provider_id" json:"shipping_provider_id"`
	AssignedBy         string    `gorm:"column:assigned_by" json:"assigned_by"`
	LabelUrl           string    `gorm:"column:label_url" json:"label_url"`
	CreateBy           string    `gorm:"column:create_by" json:"create_by"`
	CreateDtm          time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateBy           string    `gorm:"column:update_by" json:"update_by"`
	UpdateDtm          time.Time `gorm:"column:update_dtm" json:"update_dtm"`
}

func (Shipment) TableName() string {
	return "shipment"
}

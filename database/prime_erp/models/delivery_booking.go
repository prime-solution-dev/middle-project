package models

import (
	"time"
)

type DeliveryBooking struct {
	Id                     string    `gorm:"column:id;primaryKey" json:"id"`
	DeliveryCode           string    `gorm:"column:delivery_code" json:"delivery_code"`
	CompanyCode            string    `gorm:"column:company_code" json:"company_code"`
	SiteCode               string    `gorm:"column:site_code" json:"site_code"`
	DeliveryMethod         string    `gorm:"column:delivery_method" json:"delivery_method"`
	DocumentRef            string    `gorm:"column:document_ref" json:"document_ref"`
	CustomerCode           string    `gorm:"column:customer_code" json:"customer_code"`
	ShipToAddress          string    `gorm:"column:ship_to_address" json:"ship_to_address"`
	DeliveryDate           time.Time `gorm:"column:delivery_date" json:"delivery_date"`
	DeliveryTimeCode       string    `gorm:"column:delivery_time_code" json:"delivery_time_code"`
	LicensePlate           string    `gorm:"column:license_plate" json:"license_plate"`
	ContactName            string    `gorm:"column:contact_name" json:"contact_name"`
	Tel                    string    `gorm:"column:tel" json:"tel"`
	TotalWeight            float64   `gorm:"column:total_weight" json:"total_weight"`
	Status                 string    `gorm:"column:status" json:"status"`
	Remark                 string    `gorm:"column:remark" json:"remark"`
	CreateDate             time.Time `gorm:"column:create_date" json:"create_date"`
	CreateBy               string    `gorm:"column:create_by" json:"create_by"`
	UpdateDate             time.Time `gorm:"column:update_date" json:"update_date"`
	UpdateBy               string    `gorm:"column:update_by" json:"update_by"`
	BookingSlotType        string    `gorm:"column:booking_slot_type" json:"booking_slot_type"`
	StatusApproveGi        string    `gorm:"column:status_approve_gi" json:"status_approve_gi"`
	ExternalId             string    `gorm:"column:external_id" json:"external_id"`
	ExternalDocumentNumber string    `gorm:"column:external_document_number" json:"external_document_number"`
}

func (DeliveryBooking) TableName() string {
	return "delivery_booking"
}

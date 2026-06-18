package models

import (
	"time"
)

type OrderPackageItem struct {
	Id string `gorm:"column:id" json:"id"`
	PackageId string `gorm:"column:package_id" json:"package_id"`
	ProductCode string `gorm:"column:product_code" json:"product_code"`
	Qty float64 `gorm:"column:qty" json:"qty"`
	UnitCode string `gorm:"column:unit_code" json:"unit_code"`
	DocumentRefItem string `gorm:"column:document_ref_item" json:"document_ref_item"`
	CreateBy string `gorm:"column:create_by" json:"create_by"`
	CreateDtm time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateBy string `gorm:"column:update_by" json:"update_by"`
	UpdateDtm time.Time `gorm:"column:update_dtm" json:"update_dtm"`
}

func (OrderPackageItem) TableName() string {
	return "order_package_item"
}

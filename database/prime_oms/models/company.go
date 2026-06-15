package models

import (
	"time"
	"gorm.io/datatypes"
)

type Company struct {
	Id string `gorm:"column:id" json:"id"`
	TenantId string `gorm:"column:tenant_id" json:"tenant_id"`
	CompanyCode string `gorm:"column:company_code" json:"company_code"`
	CompanyName string `gorm:"column:company_name" json:"company_name"`
	CreateBy string `gorm:"column:create_by" json:"create_by"`
	CreateDtm time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateBy string `gorm:"column:update_by" json:"update_by"`
	UpdateDtm time.Time `gorm:"column:update_dtm" json:"update_dtm"`
	IsActive bool `gorm:"column:is_active" json:"is_active"`
	Contact string `gorm:"column:contact" json:"contact"`
	Email string `gorm:"column:email" json:"email"`
	Phone string `gorm:"column:phone" json:"phone"`
	IsWeight bool `gorm:"column:is_weight" json:"is_weight"`
	IsPrimeErpRef bool `gorm:"column:is_prime_erp_ref" json:"is_prime_erp_ref"`
	TaxId string `gorm:"column:tax_id" json:"tax_id"`
	IsIbItemFromPo bool `gorm:"column:is_ib_item_from_po" json:"is_ib_item_from_po"`
}

func (Company) TableName() string {
	return "company"
}

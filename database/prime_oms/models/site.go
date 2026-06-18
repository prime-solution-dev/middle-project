package models

import (
	"time"
)

type Site struct {
	Id        string    `gorm:"column:id" json:"id"`
	TenantId  string    `gorm:"column:tenant_id" json:"tenant_id"`
	CompanyId string    `gorm:"column:company_id" json:"company_id"`
	SiteCode  string    `gorm:"column:site_code" json:"site_code"`
	SiteName  string    `gorm:"column:site_name" json:"site_name"`
	CreateBy  string    `gorm:"column:create_by" json:"create_by"`
	CreateDtm time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateBy  string    `gorm:"column:update_by" json:"update_by"`
	UpdateDtm time.Time `gorm:"column:update_dtm" json:"update_dtm"`
	IsActive  bool      `gorm:"column:is_active" json:"is_active"`
	Contact   string    `gorm:"column:contact" json:"contact"`
	Email     string    `gorm:"column:email" json:"email"`
	Phone     string    `gorm:"column:phone" json:"phone"`
}

func (Site) TableName() string {
	return "site"
}

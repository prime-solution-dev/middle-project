package models

import (
	"time"
)

type Purchase struct {
	Id                          string    `gorm:"column:id;primaryKey" json:"id"`
	PurchaseCode                string    `gorm:"column:purchase_code" json:"purchase_code"`
	PurchaseType                string    `gorm:"column:purchase_type" json:"purchase_type"`
	CompanyCode                 string    `gorm:"column:company_code" json:"company_code"`
	SiteCode                    string    `gorm:"column:site_code" json:"site_code"`
	DocRefType                  string    `gorm:"column:doc_ref_type" json:"doc_ref_type"`
	DocRef                      string    `gorm:"column:doc_ref" json:"doc_ref"`
	SupplierCode                string    `gorm:"column:supplier_code" json:"supplier_code"`
	DeliveryDate                time.Time `gorm:"column:delivery_date" json:"delivery_date"`
	DeliveryAddress             string    `gorm:"column:delivery_address" json:"delivery_address"`
	Status                      string    `gorm:"column:status" json:"status"`
	TotalAmount                 float64   `gorm:"column:total_amount" json:"total_amount"`
	TotalWeight                 float64   `gorm:"column:total_weight" json:"total_weight"`
	TotalDiscount               float64   `gorm:"column:total_discount" json:"total_discount"`
	TotalVat                    float64   `gorm:"column:total_vat" json:"total_vat"`
	SubtotalExclVat             float64   `gorm:"column:subtotal_excl_vat" json:"subtotal_excl_vat"`
	IsApproved                  bool      `gorm:"column:is_approved" json:"is_approved"`
	StatusApprove               string    `gorm:"column:status_approve" json:"status_approve"`
	Remark                      string    `gorm:"column:remark" json:"remark"`
	CreateBy                    string    `gorm:"column:create_by" json:"create_by"`
	CreateDtm                   time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateBy                    string    `gorm:"column:update_by" json:"update_by"`
	UpdateDtm                   time.Time `gorm:"column:update_dtm" json:"update_dtm"`
	StatusPayment               string    `gorm:"column:status_payment" json:"status_payment"`
	SubtotalExclDiscountExclVat float64   `gorm:"column:subtotal_excl_discount_excl_vat" json:"subtotal_excl_discount_excl_vat"`
	TradingRef                  string    `gorm:"column:trading_ref" json:"trading_ref"`
	SupplierName                string    `gorm:"column:supplier_name" json:"supplier_name"`
	SupplierAddress             string    `gorm:"column:supplier_address" json:"supplier_address"`
	SupplierPhone               string    `gorm:"column:supplier_phone" json:"supplier_phone"`
	SupplierEmail               string    `gorm:"column:supplier_email" json:"supplier_email"`
	CreditTerm                  int32     `gorm:"column:credit_term" json:"credit_term"`
	UsedType                    string    `gorm:"column:used_type" json:"used_type"`
	UsedStatus                  string    `gorm:"column:used_status" json:"used_status"`
}

func (Purchase) TableName() string {
	return "purchase"
}

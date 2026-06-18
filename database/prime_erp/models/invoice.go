package models

import (
	"time"
)

type Invoice struct {
	Id                     string    `gorm:"column:id;primaryKey" json:"id"`
	InvoiceCode            string    `gorm:"column:invoice_code" json:"invoice_code"`
	DueDate                time.Time `gorm:"column:due_date" json:"due_date"`
	Status                 string    `gorm:"column:status" json:"status"`
	Remark                 string    `gorm:"column:remark" json:"remark"`
	CreateBy               string    `gorm:"column:create_by" json:"create_by"`
	CreateDtm              time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateBy               string    `gorm:"column:update_by" json:"update_by"`
	UpdateDtm              time.Time `gorm:"column:update_dtm" json:"update_dtm"`
	TotalAmount            float64   `gorm:"column:total_amount" json:"total_amount"`
	InvoiceRef             string    `gorm:"column:invoice_ref" json:"invoice_ref"`
	InvoiceType            string    `gorm:"column:invoice_type" json:"invoice_type"`
	DocumentRefType        string    `gorm:"column:document_ref_type" json:"document_ref_type"`
	DocumentRef            string    `gorm:"column:document_ref" json:"document_ref"`
	CreditTermDay          int32     `gorm:"column:credit_term_day" json:"credit_term_day"`
	PaymentDate            time.Time `gorm:"column:payment_date" json:"payment_date"`
	DocumentDate           time.Time `gorm:"column:document_date" json:"document_date"`
	TaxDate                time.Time `gorm:"column:tax_date" json:"tax_date"`
	TaxInvoice             string    `gorm:"column:tax_invoice" json:"tax_invoice"`
	PartyType              string    `gorm:"column:party_type" json:"party_type"`
	PartyCode              string    `gorm:"column:party_code" json:"party_code"`
	TotalVat               float64   `gorm:"column:total_vat" json:"total_vat"`
	PartyName              string    `gorm:"column:party_name" json:"party_name"`
	PartyBranch            string    `gorm:"column:party_branch" json:"party_branch"`
	PartyAddress           string    `gorm:"column:party_address" json:"party_address"`
	PartyEmail             string    `gorm:"column:party_email" json:"party_email"`
	PartyTel               string    `gorm:"column:party_tel" json:"party_tel"`
	PartyTaxId             string    `gorm:"column:party_tax_id" json:"party_tax_id"`
	ExternalId             string    `gorm:"column:external_id" json:"external_id"`
	PaymentMethod          string    `gorm:"column:payment_method" json:"payment_method"`
	SubmitDate             time.Time `gorm:"column:submit_date" json:"submit_date"`
	OwnerName              string    `gorm:"column:owner_name" json:"owner_name"`
	PartyDocumentRef       string    `gorm:"column:party_document_ref" json:"party_document_ref"`
	SubtotalExclVatDeposit float64   `gorm:"column:subtotal_excl_vat_deposit" json:"subtotal_excl_vat_deposit"`
	SubtotalExclVat        float64   `gorm:"column:subtotal_excl_vat" json:"subtotal_excl_vat"`
	CompanyCode            string    `gorm:"column:company_code" json:"company_code"`
	SiteCode               string    `gorm:"column:site_code" json:"site_code"`
	IsInterface            bool      `gorm:"column:is_interface" json:"is_interface"`
	InvoiceDate            time.Time `gorm:"column:invoice_date" json:"invoice_date"`
	TotalDiscount          float64   `gorm:"column:total_discount" json:"total_discount"`
}

func (Invoice) TableName() string {
	return "invoice"
}

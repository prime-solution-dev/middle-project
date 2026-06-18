package models

import (
	"time"
)

type InvoiceItem struct {
	Id                     string    `gorm:"column:id;primaryKey" json:"id"`
	InvoiceItem            string    `gorm:"column:invoice_item" json:"invoice_item"`
	InvoiceId              string    `gorm:"column:invoice_id" json:"invoice_id"`
	DocRefItem             string    `gorm:"column:doc_ref_item" json:"doc_ref_item"`
	ProductCode            string    `gorm:"column:product_code" json:"product_code"`
	Qty                    float64   `gorm:"column:qty" json:"qty"`
	UnitCode               string    `gorm:"column:unit_code" json:"unit_code"`
	PriceUnit              float64   `gorm:"column:price_unit" json:"price_unit"`
	CreateBy               string    `gorm:"column:create_by" json:"create_by"`
	CreateDtm              time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateBy               string    `gorm:"column:update_by" json:"update_by"`
	UpdateDtm              time.Time `gorm:"column:update_dtm" json:"update_dtm"`
	InvoiceQty             float64   `gorm:"column:invoice_qty" json:"invoice_qty"`
	InvoiceUnit            string    `gorm:"column:invoice_unit" json:"invoice_unit"`
	InvoiceUnitType        string    `gorm:"column:invoice_unit_type" json:"invoice_unit_type"`
	UnitUom                string    `gorm:"column:unit_uom" json:"unit_uom"`
	WeightUnit             float64   `gorm:"column:weight_unit" json:"weight_unit"`
	AvgWeightUnit          float64   `gorm:"column:avg_weight_unit" json:"avg_weight_unit"`
	InvoiceWeight          float64   `gorm:"column:invoice_weight" json:"invoice_weight"`
	TotalDiscount          float64   `gorm:"column:total_discount" json:"total_discount"`
	TotalDiscountPercent   float64   `gorm:"column:total_discount_percent" json:"total_discount_percent"`
	DocumentRefType        string    `gorm:"column:document_ref_type" json:"document_ref_type"`
	DocumentRef            string    `gorm:"column:document_ref" json:"document_ref"`
	DocumentRefItem        string    `gorm:"column:document_ref_item" json:"document_ref_item"`
	SourceType             string    `gorm:"column:source_type" json:"source_type"`
	SourceCode             string    `gorm:"column:source_code" json:"source_code"`
	SourceItem             string    `gorm:"column:source_item" json:"source_item"`
	TotalAmount            float64   `gorm:"column:total_amount" json:"total_amount"`
	TotalVat               float64   `gorm:"column:total_vat" json:"total_vat"`
	Status                 string    `gorm:"column:status" json:"status"`
	Remark                 string    `gorm:"column:remark" json:"remark"`
	Weight                 float64   `gorm:"column:weight" json:"weight"`
	ProductDescription     string    `gorm:"column:product_description" json:"product_description"`
	PriceListUnit          float64   `gorm:"column:price_list_unit" json:"price_list_unit"`
	SubtotalExclVatDeposit float64   `gorm:"column:subtotal_excl_vat_deposit" json:"subtotal_excl_vat_deposit"`
	SubtotalExclVat        float64   `gorm:"column:subtotal_excl_vat" json:"subtotal_excl_vat"`
	DocumentDate           time.Time `gorm:"column:document_date" json:"document_date"`
	ArticleCode            string    `gorm:"column:article_code" json:"article_code"`
	ProductDesc            string    `gorm:"column:product_desc" json:"product_desc"`
	ArticleType            string    `gorm:"column:article_type" json:"article_type"`
	TotalWeight            float64   `gorm:"column:total_weight" json:"total_weight"`
}

func (InvoiceItem) TableName() string {
	return "invoice_item"
}

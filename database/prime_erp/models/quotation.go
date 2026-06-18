package models

import (
	"time"
)

type Quotation struct {
	Id                          string    `gorm:"column:id;primaryKey" json:"id"`
	QuotationCode               string    `gorm:"column:quotation_code" json:"quotation_code"`
	CustomerCode                string    `gorm:"column:customer_code" json:"customer_code"`
	CustomerName                string    `gorm:"column:customer_name" json:"customer_name"`
	DeliveryDate                time.Time `gorm:"column:delivery_date" json:"delivery_date"`
	SoldToCode                  string    `gorm:"column:sold_to_code" json:"sold_to_code"`
	SoldToAddress               string    `gorm:"column:sold_to_address" json:"sold_to_address"`
	BillToCode                  string    `gorm:"column:bill_to_code" json:"bill_to_code"`
	BillToAddress               string    `gorm:"column:bill_to_address" json:"bill_to_address"`
	ShipToCode                  string    `gorm:"column:ship_to_code" json:"ship_to_code"`
	ShipToType                  string    `gorm:"column:ship_to_type" json:"ship_to_type"`
	ShipToAddress               string    `gorm:"column:ship_to_address" json:"ship_to_address"`
	DeliveryMethod              string    `gorm:"column:delivery_method" json:"delivery_method"`
	TransportCostType           string    `gorm:"column:transport_cost_type" json:"transport_cost_type"`
	PassPrice                   bool      `gorm:"column:pass_price" json:"pass_price"`
	TotalTransportCost          float64   `gorm:"column:total_transport_cost" json:"total_transport_cost"`
	TotalAmount                 float64   `gorm:"column:total_amount" json:"total_amount"`
	TotalWeight                 float64   `gorm:"column:total_weight" json:"total_weight"`
	SubtotalExclTransport       float64   `gorm:"column:subtotal_excl_transport" json:"subtotal_excl_transport"`
	SubtotalWeightExclTransport float64   `gorm:"column:subtotal_weight_excl_transport" json:"subtotal_weight_excl_transport"`
	PaymentMethod               string    `gorm:"column:payment_method" json:"payment_method"`
	PeymentTermCode             string    `gorm:"column:peyment_term_code" json:"peyment_term_code"`
	SalePersonCode              string    `gorm:"column:sale_person_code" json:"sale_person_code"`
	EffectiveDatePrice          time.Time `gorm:"column:effective_date_price" json:"effective_date_price"`
	ExpirePriceDay              int32     `gorm:"column:expire_price_day" json:"expire_price_day"`
	ExpirePriceDate             time.Time `gorm:"column:expire_price_date" json:"expire_price_date"`
	PassPriceList               string    `gorm:"column:pass_price_list" json:"pass_price_list"`
	PassAtpCheck                string    `gorm:"column:pass_atp_check" json:"pass_atp_check"`
	PassCreditLimit             string    `gorm:"column:pass_credit_limit" json:"pass_credit_limit"`
	PassPriceExpire             string    `gorm:"column:pass_price_expire" json:"pass_price_expire"`
	Status                      string    `gorm:"column:status" json:"status"`
	Remark                      string    `gorm:"column:remark" json:"remark"`
	IsApproved                  bool      `gorm:"column:is_approved" json:"is_approved"`
	StatusApprove               string    `gorm:"column:status_approve" json:"status_approve"`
	CreateDate                  time.Time `gorm:"column:create_date" json:"create_date"`
	CreateBy                    string    `gorm:"column:create_by" json:"create_by"`
	UpdateDate                  time.Time `gorm:"column:update_date" json:"update_date"`
	UpdateBy                    string    `gorm:"column:update_by" json:"update_by"`
	CompanyCode                 string    `gorm:"column:company_code" json:"company_code"`
	SiteCode                    string    `gorm:"column:site_code" json:"site_code"`
	TotalVat                    float64   `gorm:"column:total_vat" json:"total_vat"`
	TotalDiscount               float64   `gorm:"column:total_discount" json:"total_discount"`
	SubtotalExclVat             float64   `gorm:"column:subtotal_excl_vat" json:"subtotal_excl_vat"`
	TotalTransportCostVat       float64   `gorm:"column:total_transport_cost_vat" json:"total_transport_cost_vat"`
	RemarkApproval              string    `gorm:"column:remark_approval" json:"remark_approval"`
	Revision                    float64   `gorm:"column:revision" json:"revision"`
	QuotationCodeRef            string    `gorm:"column:quotation_code_ref" json:"quotation_code_ref"`
	CreditTerm                  string    `gorm:"column:credit_term" json:"credit_term"`
	PayerTerm                   string    `gorm:"column:payer_term" json:"payer_term"`
}

func (Quotation) TableName() string {
	return "quotation"
}

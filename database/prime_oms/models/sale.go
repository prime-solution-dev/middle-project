package models

import (
	"time"
)

type Sale struct {
	Id                          string    `gorm:"column:id" json:"id"`
	SaleCode                    string    `gorm:"column:sale_code" json:"sale_code"`
	CompanyCode                 string    `gorm:"column:company_code" json:"company_code"`
	SiteCode                    string    `gorm:"column:site_code" json:"site_code"`
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
	TotalTransportCost          float64   `gorm:"column:total_transport_cost" json:"total_transport_cost"`
	TotalAmount                 float64   `gorm:"column:total_amount" json:"total_amount"`
	TotalWeight                 float64   `gorm:"column:total_weight" json:"total_weight"`
	TotalVat                    float64   `gorm:"column:total_vat" json:"total_vat"`
	TotalDiscount               float64   `gorm:"column:total_discount" json:"total_discount"`
	SubtotalExclVat             float64   `gorm:"column:subtotal_excl_vat" json:"subtotal_excl_vat"`
	SubtotalExclTransport       float64   `gorm:"column:subtotal_excl_transport" json:"subtotal_excl_transport"`
	SubtotalWeightExclTransport float64   `gorm:"column:subtotal_weight_excl_transport" json:"subtotal_weight_excl_transport"`
	PaymentMethod               string    `gorm:"column:payment_method" json:"payment_method"`
	PaymentTermCode             string    `gorm:"column:payment_term_code" json:"payment_term_code"`
	SalePersonCode              string    `gorm:"column:sale_person_code" json:"sale_person_code"`
	EffectiveDatePrice          time.Time `gorm:"column:effective_date_price" json:"effective_date_price"`
	ExpirePriceDay              int32     `gorm:"column:expire_price_day" json:"expire_price_day"`
	ExpirePriceDate             time.Time `gorm:"column:expire_price_date" json:"expire_price_date"`
	PassPriceList               string    `gorm:"column:pass_price_list" json:"pass_price_list"`
	PassAtpCheck                string    `gorm:"column:pass_atp_check" json:"pass_atp_check"`
	PassCreditLimit             string    `gorm:"column:pass_credit_limit" json:"pass_credit_limit"`
	PassPriceExpire             string    `gorm:"column:pass_price_expire" json:"pass_price_expire"`
	Status                      string    `gorm:"column:status" json:"status"`
	StatusPaymentAction         string    `gorm:"column:status_payment_action" json:"status_payment_action"`
	StatusPayment               string    `gorm:"column:status_payment" json:"status_payment"`
	Remark                      string    `gorm:"column:remark" json:"remark"`
	RemarkApproval              string    `gorm:"column:remark_approval" json:"remark_approval"`
	IsApproved                  bool      `gorm:"column:is_approved" json:"is_approved"`
	StatusApprove               string    `gorm:"column:status_approve" json:"status_approve"`
	CreateDate                  time.Time `gorm:"column:create_date" json:"create_date"`
	CreateBy                    string    `gorm:"column:create_by" json:"create_by"`
	UpdateDate                  time.Time `gorm:"column:update_date" json:"update_date"`
	UpdateBy                    string    `gorm:"column:update_by" json:"update_by"`
	QuotationId                 string    `gorm:"column:quotation_id" json:"quotation_id"`
	Source                      string    `gorm:"column:source" json:"source"`
	InterfacePlatform           string    `gorm:"column:interface_platform" json:"interface_platform"`
	CustomerOrderRef            string    `gorm:"column:customer_order_ref" json:"customer_order_ref"`
	QuotationNo                 string    `gorm:"column:quotation_no" json:"quotation_no"`
	PickingType                 string    `gorm:"column:picking_type" json:"picking_type"`
	ShipDate                    time.Time `gorm:"column:ship_date" json:"ship_date"`
	TransportZone               string    `gorm:"column:transport_zone" json:"transport_zone"`
	TransportVendor             string    `gorm:"column:transport_vendor" json:"transport_vendor"`
	ShippingLeadTime            int32     `gorm:"column:shipping_lead_time" json:"shipping_lead_time"`
	MinimumShelfLife            int32     `gorm:"column:minimum_shelf_life" json:"minimum_shelf_life"`
	InvoiceUom                  string    `gorm:"column:invoice_uom" json:"invoice_uom"`
	SalesChannel                string    `gorm:"column:sales_channel" json:"sales_channel"`
	BatchCannotReverse          bool      `gorm:"column:batch_cannot_reverse" json:"batch_cannot_reverse"`
	DocumentRef                 string    `gorm:"column:document_ref" json:"document_ref"`
	DocumentRefType             string    `gorm:"column:document_ref_type" json:"document_ref_type"`
	ExternalId                  string    `gorm:"column:external_id" json:"external_id"`
	OrderType                   string    `gorm:"column:order_type" json:"order_type"`
	ReasonCode                  string    `gorm:"column:reason_code" json:"reason_code"`
	CompletedDate               time.Time `gorm:"column:completed_date" json:"completed_date"`
	PlanReceiveDate             time.Time `gorm:"column:plan_receive_date" json:"plan_receive_date"`
	DocumentRef2                string    `gorm:"column:document_ref2" json:"document_ref2"`
	PassTradeDiscount           string    `gorm:"column:pass_trade_discount" json:"pass_trade_discount"`
	PassCommercialPromo         string    `gorm:"column:pass_commercial_promo" json:"pass_commercial_promo"`
	ExternalCode                string    `gorm:"column:external_code" json:"external_code"`
	FullFillCompanyCode         string    `gorm:"column:full_fill_company_code" json:"full_fill_company_code"`
	FullFillSiteCode            string    `gorm:"column:full_fill_site_code" json:"full_fill_site_code"`
	AllowEditItem               bool      `gorm:"column:allow_edit_item" json:"allow_edit_item"`
	AllowDeleteItem             bool      `gorm:"column:allow_delete_item" json:"allow_delete_item"`
	AllowAddItem                bool      `gorm:"column:allow_add_item" json:"allow_add_item"`
	RefSaleCode                 string    `gorm:"column:ref_sale_code" json:"ref_sale_code"`
	InvoiceNo                   string    `gorm:"column:invoice_no" json:"invoice_no"`
	PaymentGateway              string    `gorm:"column:payment_gateway" json:"payment_gateway"`
	PhoneNo                     string    `gorm:"column:phone_no" json:"phone_no"`
}

func (Sale) TableName() string {
	return "sale"
}

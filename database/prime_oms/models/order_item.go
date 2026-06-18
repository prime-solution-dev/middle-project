package models

import (
	"time"
)

type OrderItem struct {
	Id                            string    `gorm:"column:id" json:"id"`
	OrderId                       string    `gorm:"column:order_id" json:"order_id"`
	OrderItem                     string    `gorm:"column:order_item" json:"order_item"`
	DocumentRefItem               string    `gorm:"column:document_ref_item" json:"document_ref_item"`
	ProductCode                   string    `gorm:"column:product_code" json:"product_code"`
	ProductType                   string    `gorm:"column:product_type" json:"product_type"`
	InterfaceOrderQty             float64   `gorm:"column:interface_order_qty" json:"interface_order_qty"`
	Qty                           float64   `gorm:"column:qty" json:"qty"`
	UnitCode                      string    `gorm:"column:unit_code" json:"unit_code"`
	IsFocGwp                      bool      `gorm:"column:is_foc_gwp" json:"is_foc_gwp"`
	WarehouseCode                 string    `gorm:"column:warehouse_code" json:"warehouse_code"`
	BatchNo                       string    `gorm:"column:batch_no" json:"batch_no"`
	SerialCode                    string    `gorm:"column:serial_code" json:"serial_code"`
	Remark                        string    `gorm:"column:remark" json:"remark"`
	Status                        string    `gorm:"column:status" json:"status"`
	CreateBy                      string    `gorm:"column:create_by" json:"create_by"`
	CreateDtm                     time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateBy                      string    `gorm:"column:update_by" json:"update_by"`
	UpdateDtm                     time.Time `gorm:"column:update_dtm" json:"update_dtm"`
	ExternalItemRef               string    `gorm:"column:external_item_ref" json:"external_item_ref"`
	Code                          string    `gorm:"column:code" json:"code"`
	Price                         float64   `gorm:"column:price" json:"price"`
	TaxAmount                     float64   `gorm:"column:tax_amount" json:"tax_amount"`
	VoucherPlatform               float64   `gorm:"column:voucher_platform" json:"voucher_platform"`
	VoucherSeller                 float64   `gorm:"column:voucher_seller" json:"voucher_seller"`
	ShippingFeeOriginal           float64   `gorm:"column:shipping_fee_original" json:"shipping_fee_original"`
	ShippingFeeDiscountSeller     float64   `gorm:"column:shipping_fee_discount_seller" json:"shipping_fee_discount_seller"`
	ShippingFeeDiscountPlatform   float64   `gorm:"column:shipping_fee_discount_platform" json:"shipping_fee_discount_platform"`
	DeliveryOption                bool      `gorm:"column:delivery_option" json:"delivery_option"`
	VoucherAmount                 float64   `gorm:"column:voucher_amount" json:"voucher_amount"`
	Reason                        string    `gorm:"column:reason" json:"reason"`
	ErrorMessage                  string    `gorm:"column:error_message" json:"error_message"`
	ShipingMethod                 string    `gorm:"column:shiping_method" json:"shiping_method"`
	ShippingCourier               string    `gorm:"column:shipping_courier" json:"shipping_courier"`
	ShippingType                  string    `gorm:"column:shipping_type" json:"shipping_type"`
	Currency                      string    `gorm:"column:currency" json:"currency"`
	IsCod                         bool      `gorm:"column:is_cod" json:"is_cod"`
	PriceUnit                     float64   `gorm:"column:price_unit" json:"price_unit"`
	InterfaceQty                  float64   `gorm:"column:interface_qty" json:"interface_qty"`
	VatAmount                     float64   `gorm:"column:vat_amount" json:"vat_amount"`
	TotalVoucher                  float64   `gorm:"column:total_voucher" json:"total_voucher"`
	ShippingMethod                string    `gorm:"column:shipping_method" json:"shipping_method"`
	PlatformCarrierId             string    `gorm:"column:platform_carrier_id" json:"platform_carrier_id"`
	ShippingFee                   float64   `gorm:"column:shipping_fee" json:"shipping_fee"`
	ProductDesc                   string    `gorm:"column:product_desc" json:"product_desc"`
	ShippingDesc                  string    `gorm:"column:shipping_desc" json:"shipping_desc"`
	RtsDtm                        time.Time `gorm:"column:rts_dtm" json:"rts_dtm"`
	TotalAmount                   float64   `gorm:"column:total_amount" json:"total_amount"`
	UnitPriceWithVat              float64   `gorm:"column:unit_price_with_vat" json:"unit_price_with_vat"`
	UnitPriceWithoutVat           float64   `gorm:"column:unit_price_without_vat" json:"unit_price_without_vat"`
	TotalLineTax                  float64   `gorm:"column:total_line_tax" json:"total_line_tax"`
	SellerBundleDiscountWithVat   float64   `gorm:"column:seller_bundle_discount_with_vat" json:"seller_bundle_discount_with_vat"`
	SellerCoFundWithVat           float64   `gorm:"column:seller_co_fund_with_vat" json:"seller_co_fund_with_vat"`
	TotalSellerDiscountWithoutVat float64   `gorm:"column:total_seller_discount_without_vat" json:"total_seller_discount_without_vat"`
	PlatformRebateWithVat         float64   `gorm:"column:platform_rebate_with_vat" json:"platform_rebate_with_vat"`
	PackageId                     string    `gorm:"column:package_id" json:"package_id"`
	PriceUnitVat                  float64   `gorm:"column:price_unit_vat" json:"price_unit_vat"`
	DiscountUnitVat               float64   `gorm:"column:discount_unit_vat" json:"discount_unit_vat"`
	TotalDiscountVat              float64   `gorm:"column:total_discount_vat" json:"total_discount_vat"`
	SubtotalAmountVat             float64   `gorm:"column:subtotal_amount_vat" json:"subtotal_amount_vat"`
	SubtotalAmount                float64   `gorm:"column:subtotal_amount" json:"subtotal_amount"`
	TotalAmound                   float64   `gorm:"column:total_amound" json:"total_amound"`
	ItemType                      string    `gorm:"column:item_type" json:"item_type"`
}

func (OrderItem) TableName() string {
	return "order_item"
}

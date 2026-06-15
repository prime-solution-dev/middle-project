package models

import (
	"time"
	"gorm.io/datatypes"
)

type Shop struct {
	Id string `gorm:"column:id" json:"id"`
	ShopCode string `gorm:"column:shop_code" json:"shop_code"`
	ShopName string `gorm:"column:shop_name" json:"shop_name"`
	Email datatypes.JSON `gorm:"column:email" json:"email"`
	PlatformCode string `gorm:"column:platform_code" json:"platform_code"`
	CreateBy string `gorm:"column:create_by" json:"create_by"`
	CreateDtm time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateBy string `gorm:"column:update_by" json:"update_by"`
	UpdateDtm time.Time `gorm:"column:update_dtm" json:"update_dtm"`
	IsAutoReconectAuth bool `gorm:"column:is_auto_reconect_auth" json:"is_auto_reconect_auth"`
	IsActive bool `gorm:"column:is_active" json:"is_active"`
	ApiKey string `gorm:"column:api_key" json:"api_key"`
	AccessToken string `gorm:"column:access_token" json:"access_token"`
	RefreshToken string `gorm:"column:refresh_token" json:"refresh_token"`
	IsAutoReconnectAuth bool `gorm:"column:is_auto_reconnect_auth" json:"is_auto_reconnect_auth"`
	PlatformShopId string `gorm:"column:platform_shop_id" json:"platform_shop_id"`
	ExternalShopId string `gorm:"column:external_shop_id" json:"external_shop_id"`
	AutoDownloadDataFlg bool `gorm:"column:auto_download_data_flg" json:"auto_download_data_flg"`
	LastSyncedDtm time.Time `gorm:"column:last_synced_dtm" json:"last_synced_dtm"`
	SyncStatusEnabled bool `gorm:"column:sync_status_enabled" json:"sync_status_enabled"`
	IsSelfFulfilled bool `gorm:"column:is_self_fulfilled" json:"is_self_fulfilled"`
	IsDefaultCourier string `gorm:"column:is_default_courier" json:"is_default_courier"`
	CustomerCode string `gorm:"column:customer_code" json:"customer_code"`
	FullFillCompanyCode string `gorm:"column:full_fill_company_code" json:"full_fill_company_code"`
	FullFillSiteCode string `gorm:"column:full_fill_site_code" json:"full_fill_site_code"`
	SiteCode string `gorm:"column:site_code" json:"site_code"`
	CompanyCode string `gorm:"column:company_code" json:"company_code"`
	TokenExpireDtm time.Time `gorm:"column:token_expire_dtm" json:"token_expire_dtm"`
	TokenRefreshExpireDtm time.Time `gorm:"column:token_refresh_expire_dtm" json:"token_refresh_expire_dtm"`
	SalesChannelCode string `gorm:"column:sales_channel_code" json:"sales_channel_code"`
	ShopShortName string `gorm:"column:shop_short_name" json:"shop_short_name"`
	CustomerCodeId string `gorm:"column:customer_code_id" json:"customer_code_id"`
	SupplierId string `gorm:"column:supplier_id" json:"supplier_id"`
	BuyerId string `gorm:"column:buyer_id" json:"buyer_id"`
}

func (Shop) TableName() string {
	return "shop"
}

package models

import (
	"time"
)

type ShopAuth struct {
	Id             string    `gorm:"column:id" json:"id"`
	ShopId         string    `gorm:"column:shop_id" json:"shop_id"`
	ExpireDtm      time.Time `gorm:"column:expire_dtm" json:"expire_dtm"`
	TokenType      string    `gorm:"column:token_type" json:"token_type"`
	AccessToken    string    `gorm:"column:access_token" json:"access_token"`
	RefreshToken   string    `gorm:"column:refresh_token" json:"refresh_token"`
	LastRefreshDtm time.Time `gorm:"column:last_refresh_dtm" json:"last_refresh_dtm"`
	Scope          string    `gorm:"column:scope" json:"scope"`
	CreateDtm      time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateDtm      time.Time `gorm:"column:update_dtm" json:"update_dtm"`
	ExpireIn       int64     `gorm:"column:expire_in" json:"expire_in"`
}

func (ShopAuth) TableName() string {
	return "shop_auth"
}

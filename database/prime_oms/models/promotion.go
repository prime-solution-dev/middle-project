package models

import (
	"time"
	"gorm.io/datatypes"
)

type Promotion struct {
	Id string `gorm:"column:id" json:"id"`
	PromoCode string `gorm:"column:promo_code" json:"promo_code"`
	PromoType string `gorm:"column:promo_type" json:"promo_type"`
	EffectiveDtm time.Time `gorm:"column:effective_dtm" json:"effective_dtm"`
	ExpireDtm time.Time `gorm:"column:expire_dtm" json:"expire_dtm"`
	Status string `gorm:"column:status" json:"status"`
	PromoName string `gorm:"column:promo_name" json:"promo_name"`
	CreateDtm time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateDtm time.Time `gorm:"column:update_dtm" json:"update_dtm"`
	CreateBy string `gorm:"column:create_by" json:"create_by"`
	UpdateBy string `gorm:"column:update_by" json:"update_by"`
	IsNotAuto bool `gorm:"column:is_not_auto" json:"is_not_auto"`
	Priority int32 `gorm:"column:priority" json:"priority"`
	Description string `gorm:"column:description" json:"description"`
	Remark string `gorm:"column:remark" json:"remark"`
	CustomerType string `gorm:"column:customer_type" json:"customer_type"`
	CustomerCode string `gorm:"column:customer_code" json:"customer_code"`
	CustomerName string `gorm:"column:customer_name" json:"customer_name"`
	CustomerGroup string `gorm:"column:customer_group" json:"customer_group"`
	FromWarehouse string `gorm:"column:from_warehouse" json:"from_warehouse"`
	PresalesPolicy string `gorm:"column:presales_policy" json:"presales_policy"`
	AllowMultipleMatch bool `gorm:"column:allow_multiple_match" json:"allow_multiple_match"`
}

func (Promotion) TableName() string {
	return "promotion"
}

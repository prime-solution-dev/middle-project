package models

import (
	"time"
	"gorm.io/datatypes"
)

type CustomerRemainingShelfLife struct {
	Id string `gorm:"column:id" json:"id"`
	CustomerCode string `gorm:"column:customer_code" json:"customer_code"`
	CustomerGroupCode string `gorm:"column:customer_group_code" json:"customer_group_code"`
	SoldToCode string `gorm:"column:sold_to_code" json:"sold_to_code"`
	ProductCode string `gorm:"column:product_code" json:"product_code"`
	MinDay int32 `gorm:"column:min_day" json:"min_day"`
	MaxDay int32 `gorm:"column:max_day" json:"max_day"`
	EffectiveDate time.Time `gorm:"column:effective_date" json:"effective_date"`
	ExpiredDate time.Time `gorm:"column:expired_date" json:"expired_date"`
	TiffaSoldToCode string `gorm:"column:tiffa_sold_to_code" json:"tiffa_sold_to_code"`
	CustomerType string `gorm:"column:customer_type" json:"customer_type"`
	ActiveFlg bool `gorm:"column:active_flg" json:"active_flg"`
	CreateBy string `gorm:"column:create_by" json:"create_by"`
	CreateDtm time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateBy string `gorm:"column:update_by" json:"update_by"`
	UpdateDtm time.Time `gorm:"column:update_dtm" json:"update_dtm"`
	DeleteBy string `gorm:"column:delete_by" json:"delete_by"`
	DeleteDtm time.Time `gorm:"column:delete_dtm" json:"delete_dtm"`
}

func (CustomerRemainingShelfLife) TableName() string {
	return "customer_remaining_shelf_life"
}

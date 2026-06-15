package models

import (
	"time"
	"gorm.io/datatypes"
)

type PriceListExtraConfig struct {
	Id string `gorm:"column:id;primaryKey" json:"id"`
	GroupCode string `gorm:"column:group_code" json:"group_code"`
	IsActive bool `gorm:"column:is_active" json:"is_active"`
	ConfigJson datatypes.JSON `gorm:"column:config_json" json:"config_json"`
	CreateBy string `gorm:"column:create_by" json:"create_by"`
	CreateDtm time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateBy string `gorm:"column:update_by" json:"update_by"`
	UpdateDtm time.Time `gorm:"column:update_dtm" json:"update_dtm"`
}

func (PriceListExtraConfig) TableName() string {
	return "price_list_extra_config"
}

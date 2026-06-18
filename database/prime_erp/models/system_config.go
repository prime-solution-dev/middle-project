package models

import (
	"gorm.io/datatypes"
)

type SystemConfig struct {
	TopicCode  string         `gorm:"column:topic_code" json:"topic_code"`
	ConfigCode string         `gorm:"column:config_code" json:"config_code"`
	TenantId   string         `gorm:"column:tenant_id" json:"tenant_id"`
	ConfigName string         `gorm:"column:config_name" json:"config_name"`
	Cond1      string         `gorm:"column:cond1" json:"cond1"`
	Cond2      string         `gorm:"column:cond2" json:"cond2"`
	Value      string         `gorm:"column:value" json:"value"`
	Sequence   int32          `gorm:"column:sequence" json:"sequence"`
	Remark     string         `gorm:"column:remark" json:"remark"`
	Json       datatypes.JSON `gorm:"column:json" json:"json"`
}

func (SystemConfig) TableName() string {
	return "system_config"
}

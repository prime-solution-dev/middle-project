package models

import "github.com/google/uuid"

type SystemConfig struct {
	ID         uuid.UUID `gorm:"column:id;primaryKey" json:"id"`
	ConfigCode *string   `gorm:"column:config_code" json:"config_code"`
	Value      *string   `gorm:"column:value" json:"value"`
}

func (SystemConfig) TableName() string {
	return "system_config"
}

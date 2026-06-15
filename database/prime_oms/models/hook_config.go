package models

import (
	"gorm.io/datatypes"
)

type HookConfig struct {
	Id         string         `gorm:"column:id" json:"id"`
	Module     string         `gorm:"column:module" json:"module"`
	Topic      string         `gorm:"column:topic" json:"topic"`
	SubTopic   string         `gorm:"column:sub_topic" json:"sub_topic"`
	HookMethod string         `gorm:"column:hook_method" json:"hook_method"`
	HookUrl    string         `gorm:"column:hook_url" json:"hook_url"`
	Header     datatypes.JSON `gorm:"column:header" json:"header"`
	Body       string         `gorm:"column:body" json:"body"`
}

func (HookConfig) TableName() string {
	return "hook_config"
}

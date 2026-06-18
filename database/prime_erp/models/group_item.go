package models

import (
	"time"
)

type GroupItem struct {
	Id        string    `gorm:"column:id;primaryKey" json:"id"`
	GroupId   string    `gorm:"column:group_id" json:"group_id"`
	ItemCode  string    `gorm:"column:item_code" json:"item_code"`
	ItemName  string    `gorm:"column:item_name" json:"item_name"`
	Value     string    `gorm:"column:value" json:"value"`
	ValueInt  float64   `gorm:"column:value_int" json:"value_int"`
	CreateDtm time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	CreateBy  string    `gorm:"column:create_by" json:"create_by"`
	UpdateDtm time.Time `gorm:"column:update_dtm" json:"update_dtm"`
	UpdateBy  string    `gorm:"column:update_by" json:"update_by"`
}

func (GroupItem) TableName() string {
	return "group_item"
}

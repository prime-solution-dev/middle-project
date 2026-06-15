package models

import (
	"time"
	"gorm.io/datatypes"
)

type UnitUom struct {
	Id string `gorm:"column:id;primaryKey" json:"id"`
	UomCode string `gorm:"column:uom_code" json:"uom_code"`
	UomName string `gorm:"column:uom_name" json:"uom_name"`
	MethodId string `gorm:"column:method_id" json:"method_id"`
}

func (UnitUom) TableName() string {
	return "unit_uom"
}

package models

import (
	"time"
	"gorm.io/datatypes"
)

type PriceListFormulas struct {
	Id string `gorm:"column:id;primaryKey" json:"id"`
	FormulaCode string `gorm:"column:formula_code" json:"formula_code"`
	Name string `gorm:"column:name" json:"name"`
	Uom string `gorm:"column:uom" json:"uom"`
	FormulaType string `gorm:"column:formula_type" json:"formula_type"`
	Expression string `gorm:"column:expression" json:"expression"`
	Params datatypes.JSON `gorm:"column:params" json:"params"`
	Rounding int32 `gorm:"column:rounding" json:"rounding"`
	CreateDtm time.Time `gorm:"column:create_dtm" json:"create_dtm"`
}

func (PriceListFormulas) TableName() string {
	return "price_list_formulas"
}

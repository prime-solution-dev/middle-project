package models

import (
	"time"
)

type SaleBom struct {
	Id            string    `gorm:"column:id" json:"id"`
	BomCode       string    `gorm:"column:bom_code" json:"bom_code"`
	BomName       string    `gorm:"column:bom_name" json:"bom_name"`
	Qty           float64   `gorm:"column:qty" json:"qty"`
	EffectiveDate time.Time `gorm:"column:effective_date" json:"effective_date"`
	ExpireDate    time.Time `gorm:"column:expire_date" json:"expire_date"`
	Remark        string    `gorm:"column:remark" json:"remark"`
	CreateBy      string    `gorm:"column:create_by" json:"create_by"`
	CreateDtm     time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateBy      string    `gorm:"column:update_by" json:"update_by"`
	UpdateDtm     time.Time `gorm:"column:update_dtm" json:"update_dtm"`
	Status        string    `gorm:"column:status" json:"status"`
	ActiveFlg     bool      `gorm:"column:active_flg" json:"active_flg"`
}

func (SaleBom) TableName() string {
	return "sale_bom"
}

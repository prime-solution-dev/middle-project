package models

import (
	"time"
	"gorm.io/datatypes"
)

type DocumentAttachment struct {
	Id string `gorm:"column:id" json:"id"`
	DocumentStepItemId string `gorm:"column:document_step_item_id" json:"document_step_item_id"`
	AttachmentType string `gorm:"column:attachment_type" json:"attachment_type"`
	ProductCode string `gorm:"column:product_code" json:"product_code"`
	Qty int32 `gorm:"column:qty" json:"qty"`
	Batch string `gorm:"column:batch" json:"batch"`
	MfgDate time.Time `gorm:"column:mfg_date" json:"mfg_date"`
	Description string `gorm:"column:description" json:"description"`
	Remark string `gorm:"column:remark" json:"remark"`
	CreateDtm time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	CreateBy string `gorm:"column:create_by" json:"create_by"`
}

func (DocumentAttachment) TableName() string {
	return "document_attachment"
}

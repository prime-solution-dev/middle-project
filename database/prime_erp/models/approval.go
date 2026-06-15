package models

import (
	"time"
	"gorm.io/datatypes"
)

type Approval struct {
	Id string `gorm:"column:id;primaryKey" json:"id"`
	ApproveCode string `gorm:"column:approve_code" json:"approve_code"`
	ApproveTopic string `gorm:"column:approve_topic" json:"approve_topic"`
	DocumentType string `gorm:"column:document_type" json:"document_type"`
	DocumentCode string `gorm:"column:document_code" json:"document_code"`
	DocumentData datatypes.JSON `gorm:"column:document_data" json:"document_data"`
	ActionDate time.Time `gorm:"column:action_date" json:"action_date"`
	Status string `gorm:"column:status" json:"status"`
	Remark string `gorm:"column:remark" json:"remark"`
	CurentStepSeq int32 `gorm:"column:curent_step_seq" json:"curent_step_seq"`
	CreateDate time.Time `gorm:"column:create_date" json:"create_date"`
	CreateBy string `gorm:"column:create_by" json:"create_by"`
	UpdateDate time.Time `gorm:"column:update_date" json:"update_date"`
	UpdateBy string `gorm:"column:update_by" json:"update_by"`
	MdItemCode string `gorm:"column:md_item_code" json:"md_item_code"`
}

func (Approval) TableName() string {
	return "approval"
}

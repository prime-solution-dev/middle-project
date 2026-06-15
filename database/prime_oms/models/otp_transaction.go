package models

import (
	"time"
	"gorm.io/datatypes"
)

type OtpTransaction struct {
	Id string `gorm:"column:id" json:"id"`
	OtpCode string `gorm:"column:otp_code" json:"otp_code"`
	DocumentStepCode string `gorm:"column:document_step_code" json:"document_step_code"`
	StepCode string `gorm:"column:step_code" json:"step_code"`
	Otp string `gorm:"column:otp" json:"otp"`
	Destination string `gorm:"column:destination" json:"destination"`
	DestinationType string `gorm:"column:destination_type" json:"destination_type"`
	DocRef string `gorm:"column:doc_ref" json:"doc_ref"`
	Status string `gorm:"column:status" json:"status"`
	ExpireDtm time.Time `gorm:"column:expire_dtm" json:"expire_dtm"`
	CreateDtm time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	CreateBy string `gorm:"column:create_by" json:"create_by"`
}

func (OtpTransaction) TableName() string {
	return "otp_transaction"
}

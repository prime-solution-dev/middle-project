package models

import (
	"time"
	"gorm.io/datatypes"
)

type CreditRequest struct {
	Id string `gorm:"column:id;primaryKey" json:"id"`
	RequestCode string `gorm:"column:request_code" json:"request_code"`
	CustomerCode string `gorm:"column:customer_code" json:"customer_code"`
	Amount float64 `gorm:"column:amount" json:"amount"`
	RequestType string `gorm:"column:request_type" json:"request_type"`
	Status string `gorm:"column:status" json:"status"`
	IsApprove bool `gorm:"column:is_approve" json:"is_approve"`
	Reason string `gorm:"column:reason" json:"reason"`
	EffectiveDtm time.Time `gorm:"column:effective_dtm" json:"effective_dtm"`
	ExpireDtm time.Time `gorm:"column:expire_dtm" json:"expire_dtm"`
	RequestDate time.Time `gorm:"column:request_date" json:"request_date"`
	ActionDate time.Time `gorm:"column:action_date" json:"action_date"`
	CreateBy string `gorm:"column:create_by" json:"create_by"`
	CreateDtm time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateBy string `gorm:"column:update_by" json:"update_by"`
	UpdateDate time.Time `gorm:"column:update_date" json:"update_date"`
	IsAction bool `gorm:"column:is_action" json:"is_action"`
	CancelReason string `gorm:"column:cancel_reason" json:"cancel_reason"`
	ApproveDate time.Time `gorm:"column:approve_date" json:"approve_date"`
}

func (CreditRequest) TableName() string {
	return "credit_request"
}

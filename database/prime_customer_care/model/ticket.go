package models

import (
	"time"

	"github.com/google/uuid"
)

type Ticket struct {
	ID            uuid.UUID `gorm:"column:id;primaryKey" json:"id"`
	TicketCode    string    `gorm:"column:ticket_code;unique" json:"ticket_code"`
	TicketChannel string    `gorm:"column:ticket_channel" json:"ticket_channel"`
	CustomerName *string `gorm:"column:customer_name" json:"customer_name"`
	Tel          *string `gorm:"column:tel" json:"tel"`
	Email        *string `gorm:"column:email" json:"email"`
	IsWrong   *bool `gorm:"column:is_wrong" json:"is_wrong"`
	IsMissing *bool `gorm:"column:is_missing" json:"is_missing"`
	StartCall *time.Time `gorm:"column:start_call" json:"start_call"`
	EndCall   *time.Time `gorm:"column:end_call" json:"end_call"`
	Status string `gorm:"column:status" json:"status"`
	CreateDate *time.Time `gorm:"column:create_date" json:"create_date"`
	CreateBy   *string    `gorm:"column:create_by" json:"create_by"`
	UpdateDate *time.Time `gorm:"column:update_date" json:"update_date"`
	UpdateBy   *string    `gorm:"column:update_by" json:"update_by"`
}

func (Ticket) TableName() string {
	return "ticket"
}
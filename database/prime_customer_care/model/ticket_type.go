package models

import (
	"time"
	"github.com/google/uuid"
)

type TicketType struct {
	ID       uuid.UUID `gorm:"column:id;primaryKey" json:"id"`
	TypeCode string    `gorm:"column:type_code;unique" json:"type_code"`
	TypeName string    `gorm:"column:type_name" json:"type_name"`
	CreateBy *string    `gorm:"column:create_by" json:"create_by"`
	CreateDtm *time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateBy *string    `gorm:"column:update_by" json:"update_by"`
	UpdateDtm *time.Time `gorm:"column:update_dtm" json:"update_dtm"`
}

func (TicketType) TableName() string {
	return "ticket_type"
}
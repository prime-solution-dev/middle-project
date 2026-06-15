package models

import (
	"time"

	"github.com/google/uuid"
)

type OpportunityTicket struct {
	ID            uuid.UUID   `gorm:"column:id;primaryKey" json:"id"`
	OpportunityID uuid.UUID   `gorm:"column:opportunity_id" json:"opportunity_id"`
	TicketCode    string      `gorm:"column:ticket_code" json:"ticket_code"`
	CreateDate    *time.Time  `gorm:"column:create_date" json:"create_date"`
	CreateBy      *string     `gorm:"column:create_by" json:"create_by"`
	UpdateDate    *time.Time  `gorm:"column:update_date" json:"update_date"`
	UpdateBy      *string     `gorm:"column:update_by" json:"update_by"`
	Opportunity   Opportunity `gorm:"foreignKey:OpportunityID"`
}

func (OpportunityTicket) TableName() string {
	return "opportunity_ticket"
}

package models

import (
	"time"

	"github.com/google/uuid"
)

type Opportunity struct {
	ID              uuid.UUID  `gorm:"column:id;primaryKey" json:"id"`
	OpportunityCode string     `gorm:"column:opportunity_code;unique" json:"opportunity_code"`
	TicketType      *string    `gorm:"column:ticket_type" json:"ticket_type"`
	OpportunityName *string    `gorm:"column:opportunity_name" json:"opportunity_name"`
	CustomerName    *string    `gorm:"column:customer_name" json:"customer_name"`
	Tel             *string    `gorm:"column:tel" json:"tel"`
	Email           *string    `gorm:"column:email" json:"email"`
	IsFollowerUp    *bool      `gorm:"column:is_follower_up" json:"is_follower_up"`
	BrandCode       *string    `gorm:"column:brand_code" json:"brand_code"`
	CsStaff         *string    `gorm:"column:cs_staff" json:"cs_staff"`
	ProductName     *string    `gorm:"column:product_name" json:"product_name"`
	OrderCode       *string    `gorm:"column:order_code" json:"order_code"`
	Status          string     `gorm:"column:status" json:"status"`
	CreateDate      *time.Time `gorm:"column:create_date" json:"create_date"`
	CreateBy        *string    `gorm:"column:create_by" json:"create_by"`
	UpdateDate      *time.Time `gorm:"column:update_date" json:"update_date"`
	UpdateBy        *string    `gorm:"column:update_by" json:"update_by"`
	CreateByChannel *string    `gorm:"column:create_by_channel" json:"create_by_channel"`

	Remarks []OpportunityRemark `gorm:"foreignKey:OpportunityID"`
	Tickets []OpportunityTicket `gorm:"foreignKey:OpportunityID"`
}

func (Opportunity) TableName() string {
	return "opportunity"
}

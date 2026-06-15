package models

import (
	"time"

	"github.com/google/uuid"
)

type OpportunityRemark struct {
	ID            uuid.UUID  `gorm:"column:id;primaryKey" json:"id"`
	OpportunityID uuid.UUID  `gorm:"column:opportunity_id" json:"opportunity_id"`
	Remark        *string    `gorm:"column:remark" json:"remark"`
	CsStaff       *string    `gorm:"column:cs_staff" json:"cs_staff"`
	CreateDate    *time.Time `gorm:"column:create_date" json:"create_date"`
	CreateBy      *string    `gorm:"column:create_by" json:"create_by"`
	UpdateDate    *time.Time `gorm:"column:update_date" json:"update_date"`
	UpdateBy      *string    `gorm:"column:update_by" json:"update_by"`
	AttachmentURL *string    `gorm:"column:attachment_url" json:"attachment_url"`

	Opportunity Opportunity `gorm:"foreignKey:OpportunityID"`

	Attachments []OpportunityRemarkAttachment `gorm:"foreignKey:OpportunityRemarkID"`
}

func (OpportunityRemark) TableName() string {
	return "opportunity_remark"
}

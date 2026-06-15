package models

import (
	"time"
	"github.com/google/uuid"
)

type OpportunityRemarkAttachment struct {
	ID                  uuid.UUID `gorm:"column:id;primaryKey" json:"id"`
	OpportunityRemarkID uuid.UUID `gorm:"column:opportunity_remark_id" json:"opportunity_remark_id"`

	FileName *string `gorm:"column:file_name" json:"file_name"`
	FilePath *string `gorm:"column:file_path" json:"file_path"`
	FileType *string `gorm:"column:file_type" json:"file_type"`
	FileSize *int64  `gorm:"column:file_size" json:"file_size"`
	CreateDate *time.Time `gorm:"column:create_date" json:"create_date"`
	CreateBy   *string    `gorm:"column:create_by" json:"create_by"`
	UpdateDate *time.Time `gorm:"column:update_date" json:"update_date"`
	UpdateBy   *string    `gorm:"column:update_by" json:"update_by"`
	Remark OpportunityRemark `gorm:"foreignKey:OpportunityRemarkID"`
}

func (OpportunityRemarkAttachment) TableName() string {
	return "opportunity_remark_attachment"
}
package models

import (
	"time"
)

type DocumentAttachmentMedia struct {
	Id                   string    `gorm:"column:id" json:"id"`
	DocumentAttachmentId string    `gorm:"column:document_attachment_id" json:"document_attachment_id"`
	MediaType            string    `gorm:"column:media_type" json:"media_type"`
	Url                  string    `gorm:"column:url" json:"url"`
	FileName             string    `gorm:"column:file_name" json:"file_name"`
	FileExtension        string    `gorm:"column:file_extension" json:"file_extension"`
	FileSize             int32     `gorm:"column:file_size" json:"file_size"`
	CreateDtm            time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	CreateBy             string    `gorm:"column:create_by" json:"create_by"`
}

func (DocumentAttachmentMedia) TableName() string {
	return "document_attachment_media"
}

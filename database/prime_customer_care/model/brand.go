package models

import "github.com/google/uuid"

type Brand struct {
	ID        uuid.UUID `gorm:"column:id;primaryKey" json:"id"`
	BrandCode string    `gorm:"column:brand_code;unique" json:"brand_code"`
	BrandName string    `gorm:"column:brand_name" json:"brand_name"`
}

func (Brand) TableName() string {
	return "brand"
}
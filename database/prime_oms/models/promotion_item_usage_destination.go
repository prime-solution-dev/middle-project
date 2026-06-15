package models

type PromotionItemUsageDestination struct {
	Id               string  `gorm:"column:id" json:"id"`
	PromoItemUsageId string  `gorm:"column:promo_item_usage_id" json:"promo_item_usage_id"`
	HierarchyType    string  `gorm:"column:hierarchy_type" json:"hierarchy_type"`
	HierarchyCode    string  `gorm:"column:hierarchy_code" json:"hierarchy_code"`
	Qty              float64 `gorm:"column:qty" json:"qty"`
}

func (PromotionItemUsageDestination) TableName() string {
	return "promotion_item_usage_destination"
}

package models

type PromotionItemDestination struct {
	Id            string  `gorm:"column:id" json:"id"`
	PromoItemId   string  `gorm:"column:promo_item_id" json:"promo_item_id"`
	HierarchyType string  `gorm:"column:hierarchy_type" json:"hierarchy_type"`
	HierarchyCode string  `gorm:"column:hierarchy_code" json:"hierarchy_code"`
	Qty           float64 `gorm:"column:qty" json:"qty"`
}

func (PromotionItemDestination) TableName() string {
	return "promotion_item_destination"
}

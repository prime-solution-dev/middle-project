package models

type PromoItemBenefitItem struct {
	Id                 string `gorm:"column:id" json:"id"`
	PromoItemBenefitId string `gorm:"column:promo_item_benefit_id" json:"promo_item_benefit_id"`
	PromoItemId        string `gorm:"column:promo_item_id" json:"promo_item_id"`
	HierarchyType      string `gorm:"column:hierarchy_type" json:"hierarchy_type"`
	HierarchyCode      string `gorm:"column:hierarchy_code" json:"hierarchy_code"`
	Qty                int32  `gorm:"column:qty" json:"qty"`
	QuotaQty           int32  `gorm:"column:quota_qty" json:"quota_qty"`
}

func (PromoItemBenefitItem) TableName() string {
	return "promo_item_benefit_item"
}

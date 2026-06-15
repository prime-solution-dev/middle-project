package models

type PromoItemBenefit struct {
	Id                       string `gorm:"column:id" json:"id"`
	PromoItemConditionItemId string `gorm:"column:promo_item_condition_item_id" json:"promo_item_condition_item_id"`
	Qty                      int32  `gorm:"column:qty" json:"qty"`
	Seq                      int32  `gorm:"column:seq" json:"seq"`
	ConditionType            string `gorm:"column:condition_type" json:"condition_type"`
}

func (PromoItemBenefit) TableName() string {
	return "promo_item_benefit"
}

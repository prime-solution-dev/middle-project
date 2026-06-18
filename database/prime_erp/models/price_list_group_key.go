package models

type PriceListGroupKey struct {
	Id               string `gorm:"column:id;primaryKey" json:"id"`
	PriceListGroupId string `gorm:"column:price_list_group_id" json:"price_list_group_id"`
	Code             string `gorm:"column:code" json:"code"`
	Value            string `gorm:"column:value" json:"value"`
	Seq              int64  `gorm:"column:seq" json:"seq"`
}

func (PriceListGroupKey) TableName() string {
	return "price_list_group_key"
}

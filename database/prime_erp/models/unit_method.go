package models

type UnitMethod struct {
	Id         string `gorm:"column:id;primaryKey" json:"id"`
	MethodCode string `gorm:"column:method_code" json:"method_code"`
	MethodName string `gorm:"column:method_name" json:"method_name"`
	UnitId     string `gorm:"column:unit_id" json:"unit_id"`
}

func (UnitMethod) TableName() string {
	return "unit_method"
}

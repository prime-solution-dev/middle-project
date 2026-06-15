package models

import (
	"time"
	"gorm.io/datatypes"
)

type Sla struct {
	Id string `gorm:"column:id" json:"id"`
	PlatformCode string `gorm:"column:platform_code" json:"platform_code"`
	StepCode string `gorm:"column:step_code" json:"step_code"`
	SlaMins int32 `gorm:"column:sla_mins" json:"sla_mins"`
	CreateBy string `gorm:"column:create_by" json:"create_by"`
	CreateDtm time.Time `gorm:"column:create_dtm" json:"create_dtm"`
	UpdateBy string `gorm:"column:update_by" json:"update_by"`
	UpdateDtm time.Time `gorm:"column:update_dtm" json:"update_dtm"`
	WorkDayMon bool `gorm:"column:work_day_mon" json:"work_day_mon"`
	WorkDayTue bool `gorm:"column:work_day_tue" json:"work_day_tue"`
	WorkDayWed bool `gorm:"column:work_day_wed" json:"work_day_wed"`
	WorkDayThu bool `gorm:"column:work_day_thu" json:"work_day_thu"`
	WorkDayFri bool `gorm:"column:work_day_fri" json:"work_day_fri"`
	WorkDaySat bool `gorm:"column:work_day_sat" json:"work_day_sat"`
	WorkDaySun bool `gorm:"column:work_day_sun" json:"work_day_sun"`
}

func (Sla) TableName() string {
	return "sla"
}

package models

type Step struct {
	Id            string `gorm:"column:id" json:"id"`
	StepCode      string `gorm:"column:step_code" json:"step_code"`
	StepName      string `gorm:"column:step_name" json:"step_name"`
	Seq           int32  `gorm:"column:seq" json:"seq"`
	ComponentCode string `gorm:"column:component_code" json:"component_code"`
	IsStart       bool   `gorm:"column:is_start" json:"is_start"`
	IsEnd         bool   `gorm:"column:is_end" json:"is_end"`
	IsActive      bool   `gorm:"column:is_active" json:"is_active"`
	IsMainStep    bool   `gorm:"column:is_main_step" json:"is_main_step"`
	WorkDayMon    bool   `gorm:"column:work_day_mon" json:"work_day_mon"`
	WorkDayTue    bool   `gorm:"column:work_day_tue" json:"work_day_tue"`
	WorkDayWed    bool   `gorm:"column:work_day_wed" json:"work_day_wed"`
	WorkDayThu    bool   `gorm:"column:work_day_thu" json:"work_day_thu"`
	WorkDayFri    bool   `gorm:"column:work_day_fri" json:"work_day_fri"`
	WorkDaySat    bool   `gorm:"column:work_day_sat" json:"work_day_sat"`
	WorkDaySun    bool   `gorm:"column:work_day_sun" json:"work_day_sun"`
}

func (Step) TableName() string {
	return "step"
}

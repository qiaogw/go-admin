package models

import (
	// "time"

	"go-admin/common/global"
	"go-admin/common/models"
)

type EnergyDimTick struct {
	models.Model

	TickCode  string `json:"tickCode" gorm:"type:varchar(255);comment:时段代码"`
	TickName  string `json:"tickName" gorm:"type:varchar(255);comment:时段名称"`
	TickStart string `json:"tickStart" gorm:"type:varchar(255);comment:时段开始时间"`
	TickEnd   string `json:"tickEnd" gorm:"type:varchar(255);comment:时段结束时间"`
	models.ModelTime
	models.ControlBy
}

func (EnergyDimTick) TableName() string {
	return global.TablePrefix + "energy_dim_tick"
}

func (EnergyDimTick) TableTitle() string {
	return "时段维"
}

func (e *EnergyDimTick) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *EnergyDimTick) GetId() interface{} {
	return e.Id
}

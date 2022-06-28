package models

import (
	"go-admin/common/global"
	"go-admin/common/models"
)

type EnergyConsumption struct {
	models.Model

	ProvinceId   string `json:"provinceId" gorm:"type:int;comment:省" title:"省"`
	ProvinceName string `json:"provinceName" gorm:"type:varchar(255);comment:省名称" title:"省名称"`
	I            string `json:"i" gorm:"type:decimal(24,6);comment:0:00" title:"0:00"`
	Ii           string `json:"ii" gorm:"type:decimal(24,6);comment:1:00" title:"1:00"`
	Iii          string `json:"iii" gorm:"type:decimal(24,6);comment:2:00" title:"2:00"`
	Iv           string `json:"iv" gorm:"type:decimal(24,6);comment:3:00" title:"3:00"`
	V            string `json:"v" gorm:"type:decimal(24,6);comment:4:00" title:"4:00"`
	Vi           string `json:"vi" gorm:"type:decimal(24,6);comment:5:00" title:"5:00"`
	Vii          string `json:"vii" gorm:"type:decimal(24,6);comment:6:00" title:"6:00"`
	Viii         string `json:"viii" gorm:"type:decimal(24,6);comment:7:00" title:"7:00"`
	Ix           string `json:"ix" gorm:"type:decimal(24,6);comment:8:00" title:"8:00"`
	X            string `json:"x" gorm:"type:decimal(24,6);comment:9:00" title:"9:00"`
	Xi           string `json:"xi" gorm:"type:decimal(24,6);comment:10:00" title:"10:00"`
	Xii          string `json:"xii" gorm:"type:decimal(24,6);comment:11:00" title:"11:00"`
	Xiii         string `json:"xiii" gorm:"type:decimal(24,6);comment:12:00" title:"12:00"`
	Xiv          string `json:"xiv" gorm:"type:decimal(24,6);comment:13:00" title:"13:00"`
	Xv           string `json:"xv" gorm:"type:decimal(24,6);comment:14:00" title:"14:00"`
	Xvi          string `json:"xvi" gorm:"type:decimal(24,6);comment:15:00" title:"15:00"`
	Xvii         string `json:"xvii" gorm:"type:decimal(24,6);comment:16:00" title:"16:00"`
	Xviii        string `json:"xviii" gorm:"type:decimal(24,6);comment:17:00" title:"17:00"`
	Xix          string `json:"xix" gorm:"type:decimal(24,6);comment:18:00" title:"18:00"`
	Xx           string `json:"xx" gorm:"type:decimal(24,6);comment:19:00" title:"19:00"`
	Xxi          string `json:"xxi" gorm:"type:decimal(24,6);comment:20:00" title:"20:00"`
	Xxii         string `json:"xxii" gorm:"type:decimal(24,6);comment:21:00" title:"21:00"`
	Xxiii        string `json:"xxiii" gorm:"type:decimal(24,6);comment:22:00" title:"22:00"`
	Xxiv         string `json:"xxiv" gorm:"type:decimal(24,6);comment:23:00" title:"23:00"`
	Total        string `json:"total" gorm:"type:decimal(24,6);comment:合计" title:"合计"`
	Year         string `json:"year" gorm:"type:year;comment:年" title:"年"`
	models.ModelTime
	models.ControlBy
}

func (EnergyConsumption) TableName() string {
	return global.TablePrefix + "energy_consumption"
}

func (e *EnergyConsumption) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *EnergyConsumption) GetId() interface{} {
	return e.Id
}

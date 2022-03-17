package models

import "go-admin/common/global"

type TbDemo struct {
	Model
	Name string `json:"name" gorm:"type:varchar(128);comment:名称"`
	ModelTime
	ControlBy
}

func (TbDemo) TableName() string {
	return global.TablePrefix + "tb_demo"
}

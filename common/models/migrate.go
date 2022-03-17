package models

import (
	"go-admin/common/global"
	"time"
)

type Migration struct {
	Version   string    `gorm:"primaryKey"`
	ApplyTime time.Time `gorm:"autoCreateTime"`
}

func (Migration) TableName() string {
	return global.TablePrefix + "sys_migration"
}

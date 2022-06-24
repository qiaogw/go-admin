package service

import (
	"errors"

    "github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/energy/models"
	"go-admin/app/energy/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type EnergyDimTick struct {
	service.Service
}

// GetPage 获取EnergyDimTick列表
func (e *EnergyDimTick) GetPage(c *dto.EnergyDimTickGetPageReq, p *actions.DataPermission, list *[]models.EnergyDimTick, count *int64) error {
	var err error
	var data models.EnergyDimTick

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("EnergyDimTickService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取EnergyDimTick对象
func (e *EnergyDimTick) Get(d *dto.EnergyDimTickGetReq, p *actions.DataPermission, model *models.EnergyDimTick) error {
	var data models.EnergyDimTick

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetEnergyDimTick error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建EnergyDimTick对象
func (e *EnergyDimTick) Insert(c *dto.EnergyDimTickInsertReq) error {
    var err error
    var data models.EnergyDimTick
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("EnergyDimTickService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改EnergyDimTick对象
func (e *EnergyDimTick) Update(c *dto.EnergyDimTickUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.EnergyDimTick{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if db.Error != nil {
        e.Log.Errorf("EnergyDimTickService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除EnergyDimTick
func (e *EnergyDimTick) Remove(d *dto.EnergyDimTickDeleteReq, p *actions.DataPermission) error {
	var data models.EnergyDimTick

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveEnergyDimTick error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
package service

import (
	"encoding/json"
	"errors"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/common/tools"
	"gorm.io/gorm"
	"io"

	"go-admin/app/energy/models"
	"go-admin/app/energy/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type EnergyConsumption struct {
	service.Service
}

// GetPage 获取EnergyConsumption列表
func (e *EnergyConsumption) GetPage(c *dto.EnergyConsumptionGetPageReq, p *actions.DataPermission, list *[]models.EnergyConsumption, count *int64) error {
	var err error
	var data models.EnergyConsumption

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("EnergyConsumptionService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取EnergyConsumption对象
func (e *EnergyConsumption) Get(d *dto.EnergyConsumptionGetReq, p *actions.DataPermission, model *models.EnergyConsumption) error {
	var data models.EnergyConsumption

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetEnergyConsumption error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建EnergyConsumption对象
func (e *EnergyConsumption) Insert(c *dto.EnergyConsumptionInsertReq) error {
	var err error
	var data models.EnergyConsumption
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("EnergyConsumptionService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改EnergyConsumption对象
func (e *EnergyConsumption) Update(c *dto.EnergyConsumptionUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.EnergyConsumption{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if db.Error != nil {
		e.Log.Errorf("EnergyConsumptionService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除EnergyConsumption
func (e *EnergyConsumption) Remove(d *dto.EnergyConsumptionDeleteReq, p *actions.DataPermission) error {
	var data models.EnergyConsumption

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveEnergyConsumption error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

// Import 导入EnergyConsumption对象
func (e *EnergyConsumption) Import(hFile io.Reader) (err error) {
	tx := e.Orm.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	var data []models.EnergyConsumption
	temp := new(models.EnergyConsumption)
	ex := new(tools.ExcelStruct)
	ex.Model = temp
	err = ex.ReadExcelIo(hFile)
	if err != nil {
		e.Log.Error(err)
		return err
	}
	err = ex.CreateMap()
	if err != nil {
		e.Log.Error(err)
		return
	}

	err = tx.Unscoped().Where("1=1").Delete(models.EnergyConsumption{}).Error
	if err != nil {
		return err
	}
	x, err := json.Marshal(ex.Info)
	if err != nil {
		e.Log.Errorf("EnergyConsumptionService Import error:%s \r\n", err)
		return err
	}
	err = json.Unmarshal(x, &data)
	if err != nil {
		e.Log.Errorf("EnergyConsumptionService Import error:%s \r\n", err)
		return err
	}
	for i := 0; i < len(data); i += 1000 {
		end := i + 1000
		if end > len(data) {
			end = len(data)
		}
		err = tx.CreateInBatches(data[i:end], len(data[i:end])).Error
		if err != nil {
			return err
		}
	}
	return nil
}

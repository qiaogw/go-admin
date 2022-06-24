package dto

import (
     
     "time"

	"go-admin/app/energy/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type EnergyDimTickGetPageReq struct {
	dto.Pagination     `search:"-"`
    TickCode string `form:"tickCode"  search:"type:exact;column:tick_code;table:energy_dim_tick" comment:"时段代码"`
    TickName string `form:"tickName"  search:"type:exact;column:tick_name;table:energy_dim_tick" comment:"时段名称"`
    EnergyDimTickOrder
}

type EnergyDimTickOrder struct {Id int `form:"idOrder"  search:"type:order;column:id;table:energy_dim_tick"`
    TickCode string `form:"tickCodeOrder"  search:"type:order;column:tick_code;table:energy_dim_tick"`
    TickName string `form:"tickNameOrder"  search:"type:order;column:tick_name;table:energy_dim_tick"`
    TickStart string `form:"tickStartOrder"  search:"type:order;column:tick_start;table:energy_dim_tick"`
    TickEnd string `form:"tickEndOrder"  search:"type:order;column:tick_end;table:energy_dim_tick"`
    CreatedAt time.Time `form:"createdAtOrder"  search:"type:order;column:created_at;table:energy_dim_tick"`
    UpdatedAt time.Time `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:energy_dim_tick"`
    DeletedAt time.Time `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:energy_dim_tick"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:energy_dim_tick"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:energy_dim_tick"`
    
}

func (m *EnergyDimTickGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type EnergyDimTickInsertReq struct {
    Id int `json:"-" comment:"主键编码"` // 主键编码
    TickCode string `json:"tickCode" comment:"时段代码"`
    TickName string `json:"tickName" comment:"时段名称"`
    TickStart string `json:"tickStart" comment:"时段开始时间"`
    TickEnd string `json:"tickEnd" comment:"时段结束时间"`
    common.ControlBy
}

func (s *EnergyDimTickInsertReq) Generate(model *models.EnergyDimTick)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.TickCode = s.TickCode
    model.TickName = s.TickName
    model.TickStart = s.TickStart
    model.TickEnd = s.TickEnd
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *EnergyDimTickInsertReq) GetId() interface{} {
	return s.Id
}

type EnergyDimTickUpdateReq struct {
    Id int `uri:"id" comment:"主键编码"` // 主键编码
    TickCode string `json:"tickCode" comment:"时段代码"`
    TickName string `json:"tickName" comment:"时段名称"`
    TickStart string `json:"tickStart" comment:"时段开始时间"`
    TickEnd string `json:"tickEnd" comment:"时段结束时间"`
    common.ControlBy
}

func (s *EnergyDimTickUpdateReq) Generate(model *models.EnergyDimTick)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.TickCode = s.TickCode
    model.TickName = s.TickName
    model.TickStart = s.TickStart
    model.TickEnd = s.TickEnd
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *EnergyDimTickUpdateReq) GetId() interface{} {
	return s.Id
}

// EnergyDimTickGetReq 功能获取请求参数
type EnergyDimTickGetReq struct {
     Id int `uri:"id"`
}
func (s *EnergyDimTickGetReq) GetId() interface{} {
	return s.Id
}

// EnergyDimTickDeleteReq 功能删除请求参数
type EnergyDimTickDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *EnergyDimTickDeleteReq) GetId() interface{} {
	return s.Ids
}
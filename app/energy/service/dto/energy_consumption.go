package dto

import (
	"time"

	"go-admin/app/energy/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type EnergyConsumptionGetPageReq struct {
	dto.Pagination `search:"-"`
	ProvinceId     string `form:"provinceId"  search:"type:exact;column:province_id;table:energy_consumption" comment:"省"`
	Year           string `form:"year"  search:"type:exact;column:year;table:energy_consumption" comment:"年"`
	EnergyConsumptionOrder
}

type EnergyConsumptionOrder struct {
	Id           int       `form:"idOrder"  search:"type:order;column:id;table:energy_consumption"`
	ProvinceId   string    `form:"provinceIdOrder"  search:"type:order;column:province_id;table:energy_consumption"`
	ProvinceName string    `form:"provinceNameOrder"  search:"type:order;column:province_name;table:energy_consumption"`
	I            string    `form:"iOrder"  search:"type:order;column:i;table:energy_consumption"`
	Ii           string    `form:"iiOrder"  search:"type:order;column:ii;table:energy_consumption"`
	Iii          string    `form:"iiiOrder"  search:"type:order;column:iii;table:energy_consumption"`
	Iv           string    `form:"ivOrder"  search:"type:order;column:iv;table:energy_consumption"`
	V            string    `form:"vOrder"  search:"type:order;column:v;table:energy_consumption"`
	Vi           string    `form:"viOrder"  search:"type:order;column:vi;table:energy_consumption"`
	Vii          string    `form:"viiOrder"  search:"type:order;column:vii;table:energy_consumption"`
	Viii         string    `form:"viiiOrder"  search:"type:order;column:viii;table:energy_consumption"`
	Ix           string    `form:"ixOrder"  search:"type:order;column:ix;table:energy_consumption"`
	X            string    `form:"xOrder"  search:"type:order;column:x;table:energy_consumption"`
	Xi           string    `form:"xiOrder"  search:"type:order;column:xi;table:energy_consumption"`
	Xii          string    `form:"xiiOrder"  search:"type:order;column:xii;table:energy_consumption"`
	Xiii         string    `form:"xiiiOrder"  search:"type:order;column:xiii;table:energy_consumption"`
	Xiv          string    `form:"xivOrder"  search:"type:order;column:xiv;table:energy_consumption"`
	Xv           string    `form:"xvOrder"  search:"type:order;column:xv;table:energy_consumption"`
	Xvi          string    `form:"xviOrder"  search:"type:order;column:xvi;table:energy_consumption"`
	Xvii         string    `form:"xviiOrder"  search:"type:order;column:xvii;table:energy_consumption"`
	Xviii        string    `form:"xviiiOrder"  search:"type:order;column:xviii;table:energy_consumption"`
	Xix          string    `form:"xixOrder"  search:"type:order;column:xix;table:energy_consumption"`
	Xx           string    `form:"xxOrder"  search:"type:order;column:xx;table:energy_consumption"`
	Xxi          string    `form:"xxiOrder"  search:"type:order;column:xxi;table:energy_consumption"`
	Xxii         string    `form:"xxiiOrder"  search:"type:order;column:xxii;table:energy_consumption"`
	Xxiii        string    `form:"xxiiiOrder"  search:"type:order;column:xxiii;table:energy_consumption"`
	Xxiv         string    `form:"xxivOrder"  search:"type:order;column:xxiv;table:energy_consumption"`
	Total        string    `form:"totalOrder"  search:"type:order;column:total;table:energy_consumption"`
	CreatedAt    time.Time `form:"createdAtOrder"  search:"type:order;column:created_at;table:energy_consumption"`
	UpdatedAt    time.Time `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:energy_consumption"`
	DeletedAt    time.Time `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:energy_consumption"`
	CreateBy     string    `form:"createByOrder"  search:"type:order;column:create_by;table:energy_consumption"`
	UpdateBy     string    `form:"updateByOrder"  search:"type:order;column:update_by;table:energy_consumption"`
	Year         string    `form:"yearOrder"  search:"type:order;column:year;table:energy_consumption"`
}

func (m *EnergyConsumptionGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type EnergyConsumptionInsertReq struct {
	Id           int    `json:"-" comment:""` //
	ProvinceId   string `json:"provinceId" comment:"省"`
	ProvinceName string `json:"provinceName" comment:"省名称"`
	I            string `json:"i" comment:"0:00"`
	Ii           string `json:"ii" comment:"1:00"`
	Iii          string `json:"iii" comment:"2:00"`
	Iv           string `json:"iv" comment:"3:00"`
	V            string `json:"v" comment:"4:00"`
	Vi           string `json:"vi" comment:"5:00"`
	Vii          string `json:"vii" comment:"6:00"`
	Viii         string `json:"viii" comment:"7:00"`
	Ix           string `json:"ix" comment:"8:00"`
	X            string `json:"x" comment:"9:00"`
	Xi           string `json:"xi" comment:"10:00"`
	Xii          string `json:"xii" comment:"11:00"`
	Xiii         string `json:"xiii" comment:"12:00"`
	Xiv          string `json:"xiv" comment:"13:00"`
	Xv           string `json:"xv" comment:"14:00"`
	Xvi          string `json:"xvi" comment:"15:00"`
	Xvii         string `json:"xvii" comment:"16:00"`
	Xviii        string `json:"xviii" comment:"17:00"`
	Xix          string `json:"xix" comment:"18:00"`
	Xx           string `json:"xx" comment:"19:00"`
	Xxi          string `json:"xxi" comment:"20:00"`
	Xxii         string `json:"xxii" comment:"21:00"`
	Xxiii        string `json:"xxiii" comment:"22:00"`
	Xxiv         string `json:"xxiv" comment:"23:00"`
	Total        string `json:"total" comment:"合计"`
	Year         string `json:"year" comment:"年"`
	common.ControlBy
}

func (s *EnergyConsumptionInsertReq) Generate(model *models.EnergyConsumption) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.ProvinceId = s.ProvinceId
	model.ProvinceName = s.ProvinceName
	model.I = s.I
	model.Ii = s.Ii
	model.Iii = s.Iii
	model.Iv = s.Iv
	model.V = s.V
	model.Vi = s.Vi
	model.Vii = s.Vii
	model.Viii = s.Viii
	model.Ix = s.Ix
	model.X = s.X
	model.Xi = s.Xi
	model.Xii = s.Xii
	model.Xiii = s.Xiii
	model.Xiv = s.Xiv
	model.Xv = s.Xv
	model.Xvi = s.Xvi
	model.Xvii = s.Xvii
	model.Xviii = s.Xviii
	model.Xix = s.Xix
	model.Xx = s.Xx
	model.Xxi = s.Xxi
	model.Xxii = s.Xxii
	model.Xxiii = s.Xxiii
	model.Xxiv = s.Xxiv
	model.Total = s.Total
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
	model.Year = s.Year
}

func (s *EnergyConsumptionInsertReq) GetId() interface{} {
	return s.Id
}

type EnergyConsumptionUpdateReq struct {
	Id           int    `uri:"id" comment:""` //
	ProvinceId   string `json:"provinceId" comment:"省"`
	ProvinceName string `json:"provinceName" comment:"省名称"`
	I            string `json:"i" comment:"0:00"`
	Ii           string `json:"ii" comment:"1:00"`
	Iii          string `json:"iii" comment:"2:00"`
	Iv           string `json:"iv" comment:"3:00"`
	V            string `json:"v" comment:"4:00"`
	Vi           string `json:"vi" comment:"5:00"`
	Vii          string `json:"vii" comment:"6:00"`
	Viii         string `json:"viii" comment:"7:00"`
	Ix           string `json:"ix" comment:"8:00"`
	X            string `json:"x" comment:"9:00"`
	Xi           string `json:"xi" comment:"10:00"`
	Xii          string `json:"xii" comment:"11:00"`
	Xiii         string `json:"xiii" comment:"12:00"`
	Xiv          string `json:"xiv" comment:"13:00"`
	Xv           string `json:"xv" comment:"14:00"`
	Xvi          string `json:"xvi" comment:"15:00"`
	Xvii         string `json:"xvii" comment:"16:00"`
	Xviii        string `json:"xviii" comment:"17:00"`
	Xix          string `json:"xix" comment:"18:00"`
	Xx           string `json:"xx" comment:"19:00"`
	Xxi          string `json:"xxi" comment:"20:00"`
	Xxii         string `json:"xxii" comment:"21:00"`
	Xxiii        string `json:"xxiii" comment:"22:00"`
	Xxiv         string `json:"xxiv" comment:"23:00"`
	Total        string `json:"total" comment:"合计"`
	Year         string `json:"year" comment:"年"`
	common.ControlBy
}

func (s *EnergyConsumptionUpdateReq) Generate(model *models.EnergyConsumption) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.ProvinceId = s.ProvinceId
	model.ProvinceName = s.ProvinceName
	model.I = s.I
	model.Ii = s.Ii
	model.Iii = s.Iii
	model.Iv = s.Iv
	model.V = s.V
	model.Vi = s.Vi
	model.Vii = s.Vii
	model.Viii = s.Viii
	model.Ix = s.Ix
	model.X = s.X
	model.Xi = s.Xi
	model.Xii = s.Xii
	model.Xiii = s.Xiii
	model.Xiv = s.Xiv
	model.Xv = s.Xv
	model.Xvi = s.Xvi
	model.Xvii = s.Xvii
	model.Xviii = s.Xviii
	model.Xix = s.Xix
	model.Xx = s.Xx
	model.Xxi = s.Xxi
	model.Xxii = s.Xxii
	model.Xxiii = s.Xxiii
	model.Xxiv = s.Xxiv
	model.Total = s.Total
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
	model.Year = s.Year
}

func (s *EnergyConsumptionUpdateReq) GetId() interface{} {
	return s.Id
}

// EnergyConsumptionGetReq 功能获取请求参数
type EnergyConsumptionGetReq struct {
	Id int `uri:"id"`
}

func (s *EnergyConsumptionGetReq) GetId() interface{} {
	return s.Id
}

// EnergyConsumptionDeleteReq 功能删除请求参数
type EnergyConsumptionDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *EnergyConsumptionDeleteReq) GetId() interface{} {
	return s.Ids
}

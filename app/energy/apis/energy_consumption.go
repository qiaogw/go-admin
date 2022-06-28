package apis

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/energy/models"
	"go-admin/app/energy/service"
	"go-admin/app/energy/service/dto"
	"go-admin/common/actions"
	"go-admin/common/tools"
)

type EnergyConsumption struct {
	api.Api
}

// GetPage 获取电力消费导入列表
// @Summary 获取电力消费导入列表
// @Description 获取电力消费导入列表
// @Tags 电力消费导入
// @Param provinceId query string false "省"
// @Param year query string false "年"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.EnergyConsumption}} "{"code": 200, "data": [...]}"
// @Router /api/v1/energy-consumption [get]
// @Security Bearer
func (e EnergyConsumption) GetPage(c *gin.Context) {
	req := dto.EnergyConsumptionGetPageReq{}
	s := service.EnergyConsumption{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.EnergyConsumption, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取电力消费导入 失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取电力消费导入
// @Summary 获取电力消费导入
// @Description 获取电力消费导入
// @Tags 电力消费导入
// @Param id path string false "id"
// @Success 200 {object} response.Response{data=models.EnergyConsumption} "{"code": 200, "data": [...]}"
// @Router /api/v1/energy-consumption/{id} [get]
// @Security Bearer
func (e EnergyConsumption) Get(c *gin.Context) {
	req := dto.EnergyConsumptionGetReq{}
	s := service.EnergyConsumption{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.EnergyConsumption

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取电力消费导入失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建电力消费导入
// @Summary 创建电力消费导入
// @Description 创建电力消费导入
// @Tags 电力消费导入
// @Accept application/json
// @Product application/json
// @Param data body dto.EnergyConsumptionInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/energy-consumption [post]
// @Security Bearer
func (e EnergyConsumption) Insert(c *gin.Context) {
	req := dto.EnergyConsumptionInsertReq{}
	s := service.EnergyConsumption{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建电力消费导入  失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改电力消费导入
// @Summary 修改电力消费导入
// @Description 修改电力消费导入
// @Tags 电力消费导入
// @Accept application/json
// @Product application/json
// @Param data body dto.EnergyConsumptionUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/energy-consumption/{id} [put]
// @Security Bearer
func (e EnergyConsumption) Update(c *gin.Context) {
	req := dto.EnergyConsumptionUpdateReq{}
	s := service.EnergyConsumption{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改电力消费导入 失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除电力消费导入
// @Summary 删除电力消费导入
// @Description 删除电力消费导入
// @Tags 电力消费导入
// @Param ids body []int false "ids"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/energy-consumption [delete]
// @Security Bearer
func (e EnergyConsumption) Delete(c *gin.Context) {
	s := service.EnergyConsumption{}
	req := dto.EnergyConsumptionDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除电力消费导入失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}

// Export 导出电力消费导入
// @Summary 导出电力消费导入
// @Description 导出电力消费导入
// @Tags 电力消费导入
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/energy-dim-tick/export [get]
// @Security Bearer
func (e EnergyConsumption) Export(c *gin.Context) {
	req := dto.EnergyConsumptionGetPageReq{}
	s := service.EnergyConsumption{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.EnergyConsumption, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("导出电力消费导入 失败，\r\n失败信息 %s", err.Error()))
		return
	}

	m := new(models.EnergyConsumption)
	tools.ExportToWeb(c, m, list, m.TableName())
}

// Import 导入电力消费导入
// @Summary 导入电力消费导入
// @Description 导入电力消费导入
// @Tags 电力消费导入
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/energy-dim-tick/import [get]
// @Security Bearer
func (e EnergyConsumption) Import(c *gin.Context) {
	files, err := c.FormFile("file")
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	//获取文件句柄
	hFile, err := files.Open()
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	s := service.EnergyConsumption{}
	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}
	err = s.Import(hFile)
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}
	e.OK(nil, "导入成功")
}

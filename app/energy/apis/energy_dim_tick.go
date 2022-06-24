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

type EnergyDimTick struct {
	api.Api
}

// GetPage 获取时段维列表
// @Summary 获取时段维列表
// @Description 获取时段维列表
// @Tags 时段维
// @Param tickCode query string false "时段代码"
// @Param tickName query string false "时段名称"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.EnergyDimTick}} "{"code": 200, "data": [...]}"
// @Router /api/v1/energy-dim-tick [get]
// @Security Bearer
func (e EnergyDimTick) GetPage(c *gin.Context) {
	req := dto.EnergyDimTickGetPageReq{}
	s := service.EnergyDimTick{}
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
	list := make([]models.EnergyDimTick, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取时段维 失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取时段维
// @Summary 获取时段维
// @Description 获取时段维
// @Tags 时段维
// @Param id path string false "id"
// @Success 200 {object} response.Response{data=models.EnergyDimTick} "{"code": 200, "data": [...]}"
// @Router /api/v1/energy-dim-tick/{id} [get]
// @Security Bearer
func (e EnergyDimTick) Get(c *gin.Context) {
	req := dto.EnergyDimTickGetReq{}
	s := service.EnergyDimTick{}
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
	var object models.EnergyDimTick

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取时段维失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建时段维
// @Summary 创建时段维
// @Description 创建时段维
// @Tags 时段维
// @Accept application/json
// @Product application/json
// @Param data body dto.EnergyDimTickInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/energy-dim-tick [post]
// @Security Bearer
func (e EnergyDimTick) Insert(c *gin.Context) {
	req := dto.EnergyDimTickInsertReq{}
	s := service.EnergyDimTick{}
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
		e.Error(500, err, fmt.Sprintf("创建时段维  失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改时段维
// @Summary 修改时段维
// @Description 修改时段维
// @Tags 时段维
// @Accept application/json
// @Product application/json
// @Param data body dto.EnergyDimTickUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/energy-dim-tick/{id} [put]
// @Security Bearer
func (e EnergyDimTick) Update(c *gin.Context) {
	req := dto.EnergyDimTickUpdateReq{}
	s := service.EnergyDimTick{}
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
		e.Error(500, err, fmt.Sprintf("修改时段维 失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除时段维
// @Summary 删除时段维
// @Description 删除时段维
// @Tags 时段维
// @Param ids body []int false "ids"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/energy-dim-tick [delete]
// @Security Bearer
func (e EnergyDimTick) Delete(c *gin.Context) {
	s := service.EnergyDimTick{}
	req := dto.EnergyDimTickDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除时段维失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}

// Export 导出时段维列表
// @Summary 导出时段维列表
// @Description 导出时段维列表
// @Tags 时段维
// @Param tickCode query string false "时段代码"
// @Param tickName query string false "时段名称"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.EnergyDimTick}} "{"code": 200, "data": [...]}"
// @Router /api/v1/energy-dim-tick/export [get]
// @Security Bearer
func (e EnergyDimTick) Export(c *gin.Context) {
	req := dto.EnergyDimTickGetPageReq{}
	s := service.EnergyDimTick{}
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
	list := make([]models.EnergyDimTick, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取时段维 失败，\r\n失败信息 %s", err.Error()))
		return
	}

	m := new(models.EnergyDimTick)
	tag := tools.GetTag(m)
	e.Logger.Debugf("tag is %+v,m.TableTitle() is %v,list is %v\n", tag, m.TableTitle(), list)
	ex := tools.NewMyExcel(m.TableTitle(), tag, list)
	e.Logger.Debugf("tag is %+v\n", tag)
	e.Logger.Debugf("ex is %+v\n", ex)
	// fmt.Printf("tag is %+v\n", tag)
	ex.ExportToWeb(c)
	// e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

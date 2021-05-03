package v1

import (
	athenaModel "athena/model"
	"gin-vue-admin/global"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags AthenaDepartment
// @Summary 创建AthenaDepartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AthenaDepartment true "创建AthenaDepartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /athenaDepartment/createAthenaDepartment [post]
func CreateAthenaDepartment(c *gin.Context) {
	var athenaDepartment athenaModel.Department
	err := c.ShouldBindJSON(&athenaDepartment)
	if err != nil {
		global.GVA_LOG.Error("[CreateAthenaDepartment]bind req fail", zap.Any("err", err))
		response.FailWithMessage("请求体不合法", c)
	}
	if err := service.CreateAthenaDepartment(c, athenaDepartment); err != nil {
		global.GVA_LOG.Error("[CreateAthenaDepartment]创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags AthenaDepartment
// @Summary 删除AthenaDepartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AthenaDepartment true "删除AthenaDepartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /athenaDepartment/deleteAthenaDepartment [delete]
func DeleteAthenaDepartment(c *gin.Context) {
	var athenaDepartment athenaModel.Department
	err := c.ShouldBindJSON(&athenaDepartment)
	if err != nil {
		global.GVA_LOG.Error("[DeleteAthenaDepartment]bind req fail", zap.Any("err", err))
		response.FailWithMessage("请求体不合法", c)
	}
	if err := service.DeleteAthenaDepartment(c, athenaDepartment); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags AthenaDepartment
// @Summary 批量删除AthenaDepartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AthenaDepartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /athenaDepartment/deleteAthenaDepartmentByIds [delete]
func DeleteAthenaDepartmentByIds(c *gin.Context) {
	var IDS athenaModel.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		global.GVA_LOG.Error("[DeleteAthenaDepartmentByIds]bind req fail", zap.Any("err", err))
		response.FailWithMessage("请求体不合法", c)
	}
	if err := service.DeleteAthenaDepartmentByIds(c, IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags AthenaDepartment
// @Summary 更新AthenaDepartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AthenaDepartment true "更新AthenaDepartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /athenaDepartment/updateAthenaDepartment [put]
func UpdateAthenaDepartment(c *gin.Context) {
	var athenaDepartment athenaModel.Department
	err := c.ShouldBindJSON(&athenaDepartment)
	if err != nil {
		global.GVA_LOG.Error("[UpdateAthenaDepartment]bind req fail", zap.Any("err", err))
		response.FailWithMessage("请求体不合法", c)
	}
	if err := service.UpdateAthenaDepartment(c, athenaDepartment); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags AthenaDepartment
// @Summary 用id查询AthenaDepartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AthenaDepartment true "用id查询AthenaDepartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /athenaDepartment/findAthenaDepartment [get]
func FindAthenaDepartment(c *gin.Context) {
	var athenaDepartment athenaModel.IdReq
	err := c.ShouldBindQuery(&athenaDepartment)
	if err != nil {
		global.GVA_LOG.Error("[FindAthenaDepartment]bind req fail", zap.Any("err", err))
		response.FailWithMessage("请求体不合法", c)
	}
	if err, reathenaDepartment := service.GetAthenaDepartment(c, athenaDepartment.Id); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reathenaDepartment": reathenaDepartment}, c)
	}
}

// @Tags AthenaDepartment
// @Summary 分页获取AthenaDepartment列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AthenaDepartmentSearch true "分页获取AthenaDepartment列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /athenaDepartment/getAthenaDepartmentList [get]
func GetAthenaDepartmentList(c *gin.Context) {
	var pageInfo request.AthenaDepartmentSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		global.GVA_LOG.Error("[GetAthenaDepartmentList]bind req fail", zap.Any("err", err))
		response.FailWithMessage("请求体不合法", c)
	}
	if err, list, total := service.GetAthenaDepartmentInfoList(c, pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

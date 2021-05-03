package v1

import (
	"gin-vue-admin/global"
    "gin-vue-admin/model"
    "gin-vue-admin/model/request"
    "gin-vue-admin/model/response"
    "gin-vue-admin/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

// @Tags AthenaDepartmentPlatform
// @Summary 创建AthenaDepartmentPlatform
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AthenaDepartmentPlatform true "创建AthenaDepartmentPlatform"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /athenaDepartmentPlatform/createAthenaDepartmentPlatform [post]
func CreateAthenaDepartmentPlatform(c *gin.Context) {
	var athenaDepartmentPlatform model.AthenaDepartmentPlatform
	err := c.ShouldBindJSON(&athenaDepartmentPlatform)
	if err!=nil {
		global.GVA_LOG.Error("[CreateAthenaDepartmentPlatform]bind req fail", zap.Any("err", err))
		response.FailWithMessage("请求体不合法", c)
	}
	if err := service.CreateAthenaDepartmentPlatform(athenaDepartmentPlatform); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags AthenaDepartmentPlatform
// @Summary 删除AthenaDepartmentPlatform
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AthenaDepartmentPlatform true "删除AthenaDepartmentPlatform"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /athenaDepartmentPlatform/deleteAthenaDepartmentPlatform [delete]
func DeleteAthenaDepartmentPlatform(c *gin.Context) {
	var athenaDepartmentPlatform model.AthenaDepartmentPlatform
	err := c.ShouldBindJSON(&athenaDepartmentPlatform)
	if err!=nil {
		global.GVA_LOG.Error("[DeleteAthenaDepartmentPlatform]bind req fail", zap.Any("err", err))
		response.FailWithMessage("请求体不合法", c)
	}
	if err := service.DeleteAthenaDepartmentPlatform(athenaDepartmentPlatform); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags AthenaDepartmentPlatform
// @Summary 批量删除AthenaDepartmentPlatform
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AthenaDepartmentPlatform"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /athenaDepartmentPlatform/deleteAthenaDepartmentPlatformByIds [delete]
func DeleteAthenaDepartmentPlatformByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err!=nil {
		global.GVA_LOG.Error("[DeleteAthenaDepartmentPlatformByIds]bind req fail", zap.Any("err", err))
		response.FailWithMessage("请求体不合法", c)
	}
	if err := service.DeleteAthenaDepartmentPlatformByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags AthenaDepartmentPlatform
// @Summary 更新AthenaDepartmentPlatform
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AthenaDepartmentPlatform true "更新AthenaDepartmentPlatform"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /athenaDepartmentPlatform/updateAthenaDepartmentPlatform [put]
func UpdateAthenaDepartmentPlatform(c *gin.Context) {
	var athenaDepartmentPlatform model.AthenaDepartmentPlatform
	err := c.ShouldBindJSON(&athenaDepartmentPlatform)
	if err!=nil {
		global.GVA_LOG.Error("[UpdateAthenaDepartmentPlatform]bind req fail", zap.Any("err", err))
		response.FailWithMessage("请求体不合法", c)
	}
	if err := service.UpdateAthenaDepartmentPlatform(athenaDepartmentPlatform); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags AthenaDepartmentPlatform
// @Summary 用id查询AthenaDepartmentPlatform
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AthenaDepartmentPlatform true "用id查询AthenaDepartmentPlatform"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /athenaDepartmentPlatform/findAthenaDepartmentPlatform [get]
func FindAthenaDepartmentPlatform(c *gin.Context) {
	var athenaDepartmentPlatform request.IdReq
	err := c.ShouldBindQuery(&athenaDepartmentPlatform)
	if err!=nil {
		global.GVA_LOG.Error("[FindAthenaDepartmentPlatform]bind req fail", zap.Any("err", err))
		response.FailWithMessage("请求体不合法", c)
	}
	if err, reathenaDepartmentPlatform := service.GetAthenaDepartmentPlatform(athenaDepartmentPlatform.Id); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reathenaDepartmentPlatform": reathenaDepartmentPlatform}, c)
	}
}

// @Tags AthenaDepartmentPlatform
// @Summary 分页获取AthenaDepartmentPlatform列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AthenaDepartmentPlatformSearch true "分页获取AthenaDepartmentPlatform列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /athenaDepartmentPlatform/getAthenaDepartmentPlatformList [get]
func GetAthenaDepartmentPlatformList(c *gin.Context) {
	var pageInfo request.AthenaDepartmentPlatformSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err!=nil {
		global.GVA_LOG.Error("[GetAthenaDepartmentPlatformList]bind req fail", zap.Any("err", err))
		response.FailWithMessage("请求体不合法", c)
	}
	if err, list, total := service.GetAthenaDepartmentPlatformInfoList(pageInfo); err != nil {
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

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

// @Tags AthenaUser
// @Summary 创建AthenaUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AthenaUser true "创建AthenaUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /athenaUser/createAthenaUser [post]
func CreateAthenaUser(c *gin.Context) {
	var athenaUser model.AthenaUser
	err := c.ShouldBindJSON(&athenaUser)
	if err!=nil {
		global.GVA_LOG.Error("[CreateAthenaUser]bind req fail", zap.Any("err", err))
		response.FailWithMessage("请求体不合法", c)
	}
	if err := service.CreateAthenaUser(athenaUser); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags AthenaUser
// @Summary 删除AthenaUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AthenaUser true "删除AthenaUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /athenaUser/deleteAthenaUser [delete]
func DeleteAthenaUser(c *gin.Context) {
	var athenaUser model.AthenaUser
	err := c.ShouldBindJSON(&athenaUser)
	if err!=nil {
		global.GVA_LOG.Error("[DeleteAthenaUser]bind req fail", zap.Any("err", err))
		response.FailWithMessage("请求体不合法", c)
	}
	if err := service.DeleteAthenaUser(athenaUser); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags AthenaUser
// @Summary 批量删除AthenaUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AthenaUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /athenaUser/deleteAthenaUserByIds [delete]
func DeleteAthenaUserByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err!=nil {
		global.GVA_LOG.Error("[DeleteAthenaUserByIds]bind req fail", zap.Any("err", err))
		response.FailWithMessage("请求体不合法", c)
	}
	if err := service.DeleteAthenaUserByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags AthenaUser
// @Summary 更新AthenaUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AthenaUser true "更新AthenaUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /athenaUser/updateAthenaUser [put]
func UpdateAthenaUser(c *gin.Context) {
	var athenaUser model.AthenaUser
	err := c.ShouldBindJSON(&athenaUser)
	if err!=nil {
		global.GVA_LOG.Error("[UpdateAthenaUser]bind req fail", zap.Any("err", err))
		response.FailWithMessage("请求体不合法", c)
	}
	if err := service.UpdateAthenaUser(athenaUser); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags AthenaUser
// @Summary 用id查询AthenaUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AthenaUser true "用id查询AthenaUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /athenaUser/findAthenaUser [get]
func FindAthenaUser(c *gin.Context) {
	var athenaUser request.IdReq
	err := c.ShouldBindQuery(&athenaUser)
	if err!=nil {
		global.GVA_LOG.Error("[FindAthenaUser]bind req fail", zap.Any("err", err))
		response.FailWithMessage("请求体不合法", c)
	}
	if err, reathenaUser := service.GetAthenaUser(athenaUser.Id); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reathenaUser": reathenaUser}, c)
	}
}

// @Tags AthenaUser
// @Summary 分页获取AthenaUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AthenaUserSearch true "分页获取AthenaUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /athenaUser/getAthenaUserList [get]
func GetAthenaUserList(c *gin.Context) {
	var pageInfo request.AthenaUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err!=nil {
		global.GVA_LOG.Error("[GetAthenaUserList]bind req fail", zap.Any("err", err))
		response.FailWithMessage("请求体不合法", c)
	}
	if err, list, total := service.GetAthenaUserInfoList(pageInfo); err != nil {
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

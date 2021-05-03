package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitAthenaDepartmentPlatformRouter(Router *gin.RouterGroup) {
	AthenaDepartmentPlatformRouter := Router.Group("athenaDepartmentPlatform").Use(middleware.OperationRecord())
	{
		AthenaDepartmentPlatformRouter.POST("createAthenaDepartmentPlatform", v1.CreateAthenaDepartmentPlatform)   // 新建AthenaDepartmentPlatform
		AthenaDepartmentPlatformRouter.DELETE("deleteAthenaDepartmentPlatform", v1.DeleteAthenaDepartmentPlatform) // 删除AthenaDepartmentPlatform
		AthenaDepartmentPlatformRouter.DELETE("deleteAthenaDepartmentPlatformByIds", v1.DeleteAthenaDepartmentPlatformByIds) // 批量删除AthenaDepartmentPlatform
		AthenaDepartmentPlatformRouter.PUT("updateAthenaDepartmentPlatform", v1.UpdateAthenaDepartmentPlatform)    // 更新AthenaDepartmentPlatform
		AthenaDepartmentPlatformRouter.GET("findAthenaDepartmentPlatform", v1.FindAthenaDepartmentPlatform)        // 根据ID获取AthenaDepartmentPlatform
		AthenaDepartmentPlatformRouter.GET("getAthenaDepartmentPlatformList", v1.GetAthenaDepartmentPlatformList)  // 获取AthenaDepartmentPlatform列表
	}
}

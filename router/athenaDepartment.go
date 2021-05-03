package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitAthenaDepartmentRouter(Router *gin.RouterGroup) {
	AthenaDepartmentRouter := Router.Group("athenaDepartment").Use(middleware.OperationRecord())
	{
		AthenaDepartmentRouter.POST("createAthenaDepartment", v1.CreateAthenaDepartment)   // 新建AthenaDepartment
		AthenaDepartmentRouter.DELETE("deleteAthenaDepartment", v1.DeleteAthenaDepartment) // 删除AthenaDepartment
		AthenaDepartmentRouter.DELETE("deleteAthenaDepartmentByIds", v1.DeleteAthenaDepartmentByIds) // 批量删除AthenaDepartment
		AthenaDepartmentRouter.PUT("updateAthenaDepartment", v1.UpdateAthenaDepartment)    // 更新AthenaDepartment
		AthenaDepartmentRouter.GET("findAthenaDepartment", v1.FindAthenaDepartment)        // 根据ID获取AthenaDepartment
		AthenaDepartmentRouter.GET("getAthenaDepartmentList", v1.GetAthenaDepartmentList)  // 获取AthenaDepartment列表
	}
}

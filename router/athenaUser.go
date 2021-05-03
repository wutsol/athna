package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitAthenaUserRouter(Router *gin.RouterGroup) {
	AthenaUserRouter := Router.Group("athenaUser").Use(middleware.OperationRecord())
	{
		AthenaUserRouter.POST("createAthenaUser", v1.CreateAthenaUser)   // 新建AthenaUser
		AthenaUserRouter.DELETE("deleteAthenaUser", v1.DeleteAthenaUser) // 删除AthenaUser
		AthenaUserRouter.DELETE("deleteAthenaUserByIds", v1.DeleteAthenaUserByIds) // 批量删除AthenaUser
		AthenaUserRouter.PUT("updateAthenaUser", v1.UpdateAthenaUser)    // 更新AthenaUser
		AthenaUserRouter.GET("findAthenaUser", v1.FindAthenaUser)        // 根据ID获取AthenaUser
		AthenaUserRouter.GET("getAthenaUserList", v1.GetAthenaUserList)  // 获取AthenaUser列表
	}
}

package test

import (
	"server/api/v1"
	"server/middleware"
	"github.com/gin-gonic/gin"
)

type LocalbillRouter struct {
}

// InitLocalbillRouter 初始化 Localbill 路由信息
func (s *LocalbillRouter) InitLocalbillRouter(Router *gin.RouterGroup) {
	lbRouter := Router.Group("lb").Use(middleware.OperationRecord())
	lbRouterWithoutRecord := Router.Group("lb")
	var lbApi = v1.ApiGroupApp.TestApiGroup.LocalbillApi
	{
		lbRouter.POST("createLocalbill", lbApi.CreateLocalbill)   // 新建Localbill
		lbRouter.DELETE("deleteLocalbill", lbApi.DeleteLocalbill) // 删除Localbill
		lbRouter.DELETE("deleteLocalbillByIds", lbApi.DeleteLocalbillByIds) // 批量删除Localbill
		lbRouter.PUT("updateLocalbill", lbApi.UpdateLocalbill)    // 更新Localbill
	}
	{
		lbRouterWithoutRecord.GET("findLocalbill", lbApi.FindLocalbill)        // 根据ID获取Localbill
		lbRouterWithoutRecord.GET("getLocalbillList", lbApi.GetLocalbillList)  // 获取Localbill列表
	}
}

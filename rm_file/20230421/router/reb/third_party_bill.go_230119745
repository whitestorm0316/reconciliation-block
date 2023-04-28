package reb

import (
	"server/api/v1"
	"server/middleware"
	"github.com/gin-gonic/gin"
)

type ThirdPartyBillRouter struct {
}

// InitThirdPartyBillRouter 初始化 ThirdPartyBill 路由信息
func (s *ThirdPartyBillRouter) InitThirdPartyBillRouter(Router *gin.RouterGroup) {
	tpbRouter := Router.Group("tpb").Use(middleware.OperationRecord())
	tpbRouterWithoutRecord := Router.Group("tpb")
	var tpbApi = v1.ApiGroupApp.RebApiGroup.ThirdPartyBillApi
	{
		tpbRouter.POST("createThirdPartyBill", tpbApi.CreateThirdPartyBill)   // 新建ThirdPartyBill
		tpbRouter.DELETE("deleteThirdPartyBill", tpbApi.DeleteThirdPartyBill) // 删除ThirdPartyBill
		tpbRouter.DELETE("deleteThirdPartyBillByIds", tpbApi.DeleteThirdPartyBillByIds) // 批量删除ThirdPartyBill
		tpbRouter.PUT("updateThirdPartyBill", tpbApi.UpdateThirdPartyBill)    // 更新ThirdPartyBill
	}
	{
		tpbRouterWithoutRecord.GET("findThirdPartyBill", tpbApi.FindThirdPartyBill)        // 根据ID获取ThirdPartyBill
		tpbRouterWithoutRecord.GET("getThirdPartyBillList", tpbApi.GetThirdPartyBillList)  // 获取ThirdPartyBill列表
	}
}

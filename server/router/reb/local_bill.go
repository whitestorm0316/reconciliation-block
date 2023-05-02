package reb

import (
	v1 "server/api/v1"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

type LocalBillRouter struct {
}

// InitLocalBillRouter 初始化 LocalBill 路由信息
func (s *LocalBillRouter) InitLocalBillRouter(Router *gin.RouterGroup) {
	lbRouter := Router.Group("lb").Use(middleware.OperationRecord())
	lbRouterWithoutRecord := Router.Group("lb")
	var lbApi = v1.ApiGroupApp.RebApiGroup.LocalBillApi
	{
		lbRouter.POST("createLocalBill", lbApi.CreateLocalBill)             // 新建LocalBill
		lbRouter.DELETE("deleteLocalBill", lbApi.DeleteLocalBill)           // 删除LocalBill
		lbRouter.DELETE("deleteLocalBillByIds", lbApi.DeleteLocalBillByIds) // 批量删除LocalBill
		lbRouter.PUT("updateLocalBill", lbApi.UpdateLocalBill)              // 更新LocalBill
	}
	{
		lbRouterWithoutRecord.GET("findLocalBill", lbApi.FindLocalBill)       // 根据ID获取LocalBill
		lbRouterWithoutRecord.GET("getLocalBillList", lbApi.GetLocalBillList) // 获取LocalBill列表
	}
}

package reb

import (
	"server/api/v1"
	"server/middleware"
	"github.com/gin-gonic/gin"
)

type ErrorBillRouter struct {
}

// InitErrorBillRouter 初始化 ErrorBill 路由信息
func (s *ErrorBillRouter) InitErrorBillRouter(Router *gin.RouterGroup) {
	ebRouter := Router.Group("eb").Use(middleware.OperationRecord())
	ebRouterWithoutRecord := Router.Group("eb")
	var ebApi = v1.ApiGroupApp.RebApiGroup.ErrorBillApi
	{
		ebRouter.POST("createErrorBill", ebApi.CreateErrorBill)   // 新建ErrorBill
		ebRouter.DELETE("deleteErrorBill", ebApi.DeleteErrorBill) // 删除ErrorBill
		ebRouter.DELETE("deleteErrorBillByIds", ebApi.DeleteErrorBillByIds) // 批量删除ErrorBill
		ebRouter.PUT("updateErrorBill", ebApi.UpdateErrorBill)    // 更新ErrorBill
	}
	{
		ebRouterWithoutRecord.GET("findErrorBill", ebApi.FindErrorBill)        // 根据ID获取ErrorBill
		ebRouterWithoutRecord.GET("getErrorBillList", ebApi.GetErrorBillList)  // 获取ErrorBill列表
	}
}

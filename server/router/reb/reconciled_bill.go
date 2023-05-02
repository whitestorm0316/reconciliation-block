package reb

import (
	v1 "server/api/v1"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

type ReconciledBillRouter struct {
}

// InitReconciledBillRouter 初始化 ReconciledBill 路由信息
func (s *ReconciledBillRouter) InitReconciledBillRouter(Router *gin.RouterGroup) {
	rbRouter := Router.Group("rb").Use(middleware.OperationRecord())
	rbRouterWithoutRecord := Router.Group("rb")
	var rbApi = v1.ApiGroupApp.RebApiGroup.ReconciledBillApi
	{
		rbRouter.POST("createReconciledBill", rbApi.CreateReconciledBill)             // 新建ReconciledBill
		rbRouter.DELETE("deleteReconciledBill", rbApi.DeleteReconciledBill)           // 删除ReconciledBill
		rbRouter.DELETE("deleteReconciledBillByIds", rbApi.DeleteReconciledBillByIds) // 批量删除ReconciledBill
		rbRouter.PUT("updateReconciledBill", rbApi.UpdateReconciledBill)              // 更新ReconciledBill
		rbRouter.POST("syncBlock", rbApi.SyncBlock)                                   // 同步至区块链
	}
	{
		rbRouterWithoutRecord.GET("findReconciledBill", rbApi.FindReconciledBill)       // 根据ID获取ReconciledBill
		rbRouterWithoutRecord.GET("getReconciledBillList", rbApi.GetReconciledBillList) // 获取ReconciledBill列表
	}
}

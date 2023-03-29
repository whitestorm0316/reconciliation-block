package bill

import (
	v1 "server/api/v1"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

type BillRouter struct{}

func (s *BillRouter) InitbillRouter(Router *gin.RouterGroup) {
	billRouter := Router.Group("bill").Use(middleware.OperationRecord())
	billRouterWithoutRecord := Router.Group("bill")
	billRouterApi := v1.ApiGroupApp.BillApiGroup.BillApi
	{
		billRouter.POST("createBill", billRouterApi.CreateBill) // 创建Api
		billRouter.POST("deleteBill", billRouterApi.DeleteBill) // 删除Api
		billRouter.POST("updateBill", billRouterApi.UpdateBill) // 更新api
	}
	{
		billRouterWithoutRecord.POST("getBillList", billRouterApi.GetBillList) // 获取所有api
	}
}

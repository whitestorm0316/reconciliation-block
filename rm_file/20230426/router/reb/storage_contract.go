package reb

import (
	"server/api/v1"
	"server/middleware"
	"github.com/gin-gonic/gin"
)

type StorageContractRouter struct {
}

// InitStorageContractRouter 初始化 StorageContract 路由信息
func (s *StorageContractRouter) InitStorageContractRouter(Router *gin.RouterGroup) {
	scRouter := Router.Group("sc").Use(middleware.OperationRecord())
	scRouterWithoutRecord := Router.Group("sc")
	var scApi = v1.ApiGroupApp.RebApiGroup.StorageContractApi
	{
		scRouter.POST("createStorageContract", scApi.CreateStorageContract)   // 新建StorageContract
		scRouter.DELETE("deleteStorageContract", scApi.DeleteStorageContract) // 删除StorageContract
		scRouter.DELETE("deleteStorageContractByIds", scApi.DeleteStorageContractByIds) // 批量删除StorageContract
		scRouter.PUT("updateStorageContract", scApi.UpdateStorageContract)    // 更新StorageContract
	}
	{
		scRouterWithoutRecord.GET("findStorageContract", scApi.FindStorageContract)        // 根据ID获取StorageContract
		scRouterWithoutRecord.GET("getStorageContractList", scApi.GetStorageContractList)  // 获取StorageContract列表
	}
}

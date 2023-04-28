package reb

import (
	"server/api/v1"
	"server/middleware"
	"github.com/gin-gonic/gin"
)

type BlockchainRouter struct {
}

// InitBlockchainRouter 初始化 Blockchain 路由信息
func (s *BlockchainRouter) InitBlockchainRouter(Router *gin.RouterGroup) {
	bcRouter := Router.Group("bc").Use(middleware.OperationRecord())
	bcRouterWithoutRecord := Router.Group("bc")
	var bcApi = v1.ApiGroupApp.RebApiGroup.BlockchainApi
	{
		bcRouter.POST("createBlockchain", bcApi.CreateBlockchain)   // 新建Blockchain
		bcRouter.DELETE("deleteBlockchain", bcApi.DeleteBlockchain) // 删除Blockchain
		bcRouter.DELETE("deleteBlockchainByIds", bcApi.DeleteBlockchainByIds) // 批量删除Blockchain
		bcRouter.PUT("updateBlockchain", bcApi.UpdateBlockchain)    // 更新Blockchain
	}
	{
		bcRouterWithoutRecord.GET("findBlockchain", bcApi.FindBlockchain)        // 根据ID获取Blockchain
		bcRouterWithoutRecord.GET("getBlockchainList", bcApi.GetBlockchainList)  // 获取Blockchain列表
	}
}

package reb

import (
	"server/global"
    "server/model/reb"
    "server/model/common/request"
    rebReq "server/model/reb/request"
    "server/model/common/response"
    "server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type BlockchainApi struct {
}

var bcService = service.ServiceGroupApp.RebServiceGroup.BlockchainService


// CreateBlockchain 创建Blockchain
// @Tags Blockchain
// @Summary 创建Blockchain
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body reb.Blockchain true "创建Blockchain"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bc/createBlockchain [post]
func (bcApi *BlockchainApi) CreateBlockchain(c *gin.Context) {
	var bc reb.Blockchain
	err := c.ShouldBindJSON(&bc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := bcService.CreateBlockchain(bc); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBlockchain 删除Blockchain
// @Tags Blockchain
// @Summary 删除Blockchain
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body reb.Blockchain true "删除Blockchain"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bc/deleteBlockchain [delete]
func (bcApi *BlockchainApi) DeleteBlockchain(c *gin.Context) {
	var bc reb.Blockchain
	err := c.ShouldBindJSON(&bc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := bcService.DeleteBlockchain(bc); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBlockchainByIds 批量删除Blockchain
// @Tags Blockchain
// @Summary 批量删除Blockchain
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Blockchain"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /bc/deleteBlockchainByIds [delete]
func (bcApi *BlockchainApi) DeleteBlockchainByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := bcService.DeleteBlockchainByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBlockchain 更新Blockchain
// @Tags Blockchain
// @Summary 更新Blockchain
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body reb.Blockchain true "更新Blockchain"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bc/updateBlockchain [put]
func (bcApi *BlockchainApi) UpdateBlockchain(c *gin.Context) {
	var bc reb.Blockchain
	err := c.ShouldBindJSON(&bc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := bcService.UpdateBlockchain(bc); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBlockchain 用id查询Blockchain
// @Tags Blockchain
// @Summary 用id查询Blockchain
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query reb.Blockchain true "用id查询Blockchain"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bc/findBlockchain [get]
func (bcApi *BlockchainApi) FindBlockchain(c *gin.Context) {
	var bc reb.Blockchain
	err := c.ShouldBindQuery(&bc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rebc, err := bcService.GetBlockchain(bc.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rebc": rebc}, c)
	}
}

// GetBlockchainList 分页获取Blockchain列表
// @Tags Blockchain
// @Summary 分页获取Blockchain列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query rebReq.BlockchainSearch true "分页获取Blockchain列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bc/getBlockchainList [get]
func (bcApi *BlockchainApi) GetBlockchainList(c *gin.Context) {
	var pageInfo rebReq.BlockchainSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := bcService.GetBlockchainInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}

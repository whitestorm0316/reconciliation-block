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

type ErrorBillApi struct {
}

var ebService = service.ServiceGroupApp.RebServiceGroup.ErrorBillService


// CreateErrorBill 创建ErrorBill
// @Tags ErrorBill
// @Summary 创建ErrorBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body reb.ErrorBill true "创建ErrorBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /eb/createErrorBill [post]
func (ebApi *ErrorBillApi) CreateErrorBill(c *gin.Context) {
	var eb reb.ErrorBill
	err := c.ShouldBindJSON(&eb)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ebService.CreateErrorBill(eb); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteErrorBill 删除ErrorBill
// @Tags ErrorBill
// @Summary 删除ErrorBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body reb.ErrorBill true "删除ErrorBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /eb/deleteErrorBill [delete]
func (ebApi *ErrorBillApi) DeleteErrorBill(c *gin.Context) {
	var eb reb.ErrorBill
	err := c.ShouldBindJSON(&eb)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ebService.DeleteErrorBill(eb); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteErrorBillByIds 批量删除ErrorBill
// @Tags ErrorBill
// @Summary 批量删除ErrorBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ErrorBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /eb/deleteErrorBillByIds [delete]
func (ebApi *ErrorBillApi) DeleteErrorBillByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ebService.DeleteErrorBillByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateErrorBill 更新ErrorBill
// @Tags ErrorBill
// @Summary 更新ErrorBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body reb.ErrorBill true "更新ErrorBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /eb/updateErrorBill [put]
func (ebApi *ErrorBillApi) UpdateErrorBill(c *gin.Context) {
	var eb reb.ErrorBill
	err := c.ShouldBindJSON(&eb)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ebService.UpdateErrorBill(eb); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindErrorBill 用id查询ErrorBill
// @Tags ErrorBill
// @Summary 用id查询ErrorBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query reb.ErrorBill true "用id查询ErrorBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /eb/findErrorBill [get]
func (ebApi *ErrorBillApi) FindErrorBill(c *gin.Context) {
	var eb reb.ErrorBill
	err := c.ShouldBindQuery(&eb)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reeb, err := ebService.GetErrorBill(eb.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reeb": reeb}, c)
	}
}

// GetErrorBillList 分页获取ErrorBill列表
// @Tags ErrorBill
// @Summary 分页获取ErrorBill列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query rebReq.ErrorBillSearch true "分页获取ErrorBill列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /eb/getErrorBillList [get]
func (ebApi *ErrorBillApi) GetErrorBillList(c *gin.Context) {
	var pageInfo rebReq.ErrorBillSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := ebService.GetErrorBillInfoList(pageInfo); err != nil {
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

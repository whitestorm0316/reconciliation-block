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

type LocalBillApi struct {
}

var lbService = service.ServiceGroupApp.RebServiceGroup.LocalBillService


// CreateLocalBill 创建LocalBill
// @Tags LocalBill
// @Summary 创建LocalBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body reb.LocalBill true "创建LocalBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lb/createLocalBill [post]
func (lbApi *LocalBillApi) CreateLocalBill(c *gin.Context) {
	var lb reb.LocalBill
	err := c.ShouldBindJSON(&lb)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := lbService.CreateLocalBill(lb); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteLocalBill 删除LocalBill
// @Tags LocalBill
// @Summary 删除LocalBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body reb.LocalBill true "删除LocalBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lb/deleteLocalBill [delete]
func (lbApi *LocalBillApi) DeleteLocalBill(c *gin.Context) {
	var lb reb.LocalBill
	err := c.ShouldBindJSON(&lb)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := lbService.DeleteLocalBill(lb); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteLocalBillByIds 批量删除LocalBill
// @Tags LocalBill
// @Summary 批量删除LocalBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LocalBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /lb/deleteLocalBillByIds [delete]
func (lbApi *LocalBillApi) DeleteLocalBillByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := lbService.DeleteLocalBillByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateLocalBill 更新LocalBill
// @Tags LocalBill
// @Summary 更新LocalBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body reb.LocalBill true "更新LocalBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lb/updateLocalBill [put]
func (lbApi *LocalBillApi) UpdateLocalBill(c *gin.Context) {
	var lb reb.LocalBill
	err := c.ShouldBindJSON(&lb)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := lbService.UpdateLocalBill(lb); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindLocalBill 用id查询LocalBill
// @Tags LocalBill
// @Summary 用id查询LocalBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query reb.LocalBill true "用id查询LocalBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lb/findLocalBill [get]
func (lbApi *LocalBillApi) FindLocalBill(c *gin.Context) {
	var lb reb.LocalBill
	err := c.ShouldBindQuery(&lb)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if relb, err := lbService.GetLocalBill(lb.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"relb": relb}, c)
	}
}

// GetLocalBillList 分页获取LocalBill列表
// @Tags LocalBill
// @Summary 分页获取LocalBill列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query rebReq.LocalBillSearch true "分页获取LocalBill列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lb/getLocalBillList [get]
func (lbApi *LocalBillApi) GetLocalBillList(c *gin.Context) {
	var pageInfo rebReq.LocalBillSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := lbService.GetLocalBillInfoList(pageInfo); err != nil {
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

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

type ThirdPartyBillApi struct {
}

var tpbService = service.ServiceGroupApp.RebServiceGroup.ThirdPartyBillService


// CreateThirdPartyBill 创建ThirdPartyBill
// @Tags ThirdPartyBill
// @Summary 创建ThirdPartyBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body reb.ThirdPartyBill true "创建ThirdPartyBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tpb/createThirdPartyBill [post]
func (tpbApi *ThirdPartyBillApi) CreateThirdPartyBill(c *gin.Context) {
	var tpb reb.ThirdPartyBill
	err := c.ShouldBindJSON(&tpb)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := tpbService.CreateThirdPartyBill(tpb); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteThirdPartyBill 删除ThirdPartyBill
// @Tags ThirdPartyBill
// @Summary 删除ThirdPartyBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body reb.ThirdPartyBill true "删除ThirdPartyBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tpb/deleteThirdPartyBill [delete]
func (tpbApi *ThirdPartyBillApi) DeleteThirdPartyBill(c *gin.Context) {
	var tpb reb.ThirdPartyBill
	err := c.ShouldBindJSON(&tpb)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := tpbService.DeleteThirdPartyBill(tpb); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteThirdPartyBillByIds 批量删除ThirdPartyBill
// @Tags ThirdPartyBill
// @Summary 批量删除ThirdPartyBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ThirdPartyBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /tpb/deleteThirdPartyBillByIds [delete]
func (tpbApi *ThirdPartyBillApi) DeleteThirdPartyBillByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := tpbService.DeleteThirdPartyBillByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateThirdPartyBill 更新ThirdPartyBill
// @Tags ThirdPartyBill
// @Summary 更新ThirdPartyBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body reb.ThirdPartyBill true "更新ThirdPartyBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tpb/updateThirdPartyBill [put]
func (tpbApi *ThirdPartyBillApi) UpdateThirdPartyBill(c *gin.Context) {
	var tpb reb.ThirdPartyBill
	err := c.ShouldBindJSON(&tpb)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := tpbService.UpdateThirdPartyBill(tpb); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindThirdPartyBill 用id查询ThirdPartyBill
// @Tags ThirdPartyBill
// @Summary 用id查询ThirdPartyBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query reb.ThirdPartyBill true "用id查询ThirdPartyBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tpb/findThirdPartyBill [get]
func (tpbApi *ThirdPartyBillApi) FindThirdPartyBill(c *gin.Context) {
	var tpb reb.ThirdPartyBill
	err := c.ShouldBindQuery(&tpb)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if retpb, err := tpbService.GetThirdPartyBill(tpb.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retpb": retpb}, c)
	}
}

// GetThirdPartyBillList 分页获取ThirdPartyBill列表
// @Tags ThirdPartyBill
// @Summary 分页获取ThirdPartyBill列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query rebReq.ThirdPartyBillSearch true "分页获取ThirdPartyBill列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tpb/getThirdPartyBillList [get]
func (tpbApi *ThirdPartyBillApi) GetThirdPartyBillList(c *gin.Context) {
	var pageInfo rebReq.ThirdPartyBillSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := tpbService.GetThirdPartyBillInfoList(pageInfo); err != nil {
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

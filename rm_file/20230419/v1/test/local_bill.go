package test

import (
	"server/global"
    "server/model/test"
    "server/model/common/request"
    testReq "server/model/test/request"
    "server/model/common/response"
    "server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type LocalbillApi struct {
}

var lbService = service.ServiceGroupApp.TestServiceGroup.LocalbillService


// CreateLocalbill 创建Localbill
// @Tags Localbill
// @Summary 创建Localbill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body test.Localbill true "创建Localbill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lb/createLocalbill [post]
func (lbApi *LocalbillApi) CreateLocalbill(c *gin.Context) {
	var lb test.Localbill
	err := c.ShouldBindJSON(&lb)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := lbService.CreateLocalbill(lb); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteLocalbill 删除Localbill
// @Tags Localbill
// @Summary 删除Localbill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body test.Localbill true "删除Localbill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lb/deleteLocalbill [delete]
func (lbApi *LocalbillApi) DeleteLocalbill(c *gin.Context) {
	var lb test.Localbill
	err := c.ShouldBindJSON(&lb)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := lbService.DeleteLocalbill(lb); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteLocalbillByIds 批量删除Localbill
// @Tags Localbill
// @Summary 批量删除Localbill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Localbill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /lb/deleteLocalbillByIds [delete]
func (lbApi *LocalbillApi) DeleteLocalbillByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := lbService.DeleteLocalbillByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateLocalbill 更新Localbill
// @Tags Localbill
// @Summary 更新Localbill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body test.Localbill true "更新Localbill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lb/updateLocalbill [put]
func (lbApi *LocalbillApi) UpdateLocalbill(c *gin.Context) {
	var lb test.Localbill
	err := c.ShouldBindJSON(&lb)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := lbService.UpdateLocalbill(lb); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindLocalbill 用id查询Localbill
// @Tags Localbill
// @Summary 用id查询Localbill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query test.Localbill true "用id查询Localbill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lb/findLocalbill [get]
func (lbApi *LocalbillApi) FindLocalbill(c *gin.Context) {
	var lb test.Localbill
	err := c.ShouldBindQuery(&lb)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if relb, err := lbService.GetLocalbill(lb.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"relb": relb}, c)
	}
}

// GetLocalbillList 分页获取Localbill列表
// @Tags Localbill
// @Summary 分页获取Localbill列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query testReq.LocalbillSearch true "分页获取Localbill列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lb/getLocalbillList [get]
func (lbApi *LocalbillApi) GetLocalbillList(c *gin.Context) {
	var pageInfo testReq.LocalbillSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := lbService.GetLocalbillInfoList(pageInfo); err != nil {
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

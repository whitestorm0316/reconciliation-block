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
    "server/utils"
)

type ReconciledBillApi struct {
}

var rbService = service.ServiceGroupApp.RebServiceGroup.ReconciledBillService


// CreateReconciledBill 创建ReconciledBill
// @Tags ReconciledBill
// @Summary 创建ReconciledBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body reb.ReconciledBill true "创建ReconciledBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rb/createReconciledBill [post]
func (rbApi *ReconciledBillApi) CreateReconciledBill(c *gin.Context) {
	var rb reb.ReconciledBill
	err := c.ShouldBindJSON(&rb)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "TransactionNumber":{utils.NotEmpty()},
        "Money":{utils.NotEmpty()},
    }
	if err := utils.Verify(rb, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := rbService.CreateReconciledBill(rb); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteReconciledBill 删除ReconciledBill
// @Tags ReconciledBill
// @Summary 删除ReconciledBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body reb.ReconciledBill true "删除ReconciledBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rb/deleteReconciledBill [delete]
func (rbApi *ReconciledBillApi) DeleteReconciledBill(c *gin.Context) {
	var rb reb.ReconciledBill
	err := c.ShouldBindJSON(&rb)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := rbService.DeleteReconciledBill(rb); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteReconciledBillByIds 批量删除ReconciledBill
// @Tags ReconciledBill
// @Summary 批量删除ReconciledBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ReconciledBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /rb/deleteReconciledBillByIds [delete]
func (rbApi *ReconciledBillApi) DeleteReconciledBillByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := rbService.DeleteReconciledBillByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateReconciledBill 更新ReconciledBill
// @Tags ReconciledBill
// @Summary 更新ReconciledBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body reb.ReconciledBill true "更新ReconciledBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /rb/updateReconciledBill [put]
func (rbApi *ReconciledBillApi) UpdateReconciledBill(c *gin.Context) {
	var rb reb.ReconciledBill
	err := c.ShouldBindJSON(&rb)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "TransactionNumber":{utils.NotEmpty()},
          "Money":{utils.NotEmpty()},
      }
    if err := utils.Verify(rb, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := rbService.UpdateReconciledBill(rb); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindReconciledBill 用id查询ReconciledBill
// @Tags ReconciledBill
// @Summary 用id查询ReconciledBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query reb.ReconciledBill true "用id查询ReconciledBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /rb/findReconciledBill [get]
func (rbApi *ReconciledBillApi) FindReconciledBill(c *gin.Context) {
	var rb reb.ReconciledBill
	err := c.ShouldBindQuery(&rb)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rerb, err := rbService.GetReconciledBill(rb.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rerb": rerb}, c)
	}
}

// GetReconciledBillList 分页获取ReconciledBill列表
// @Tags ReconciledBill
// @Summary 分页获取ReconciledBill列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query rebReq.ReconciledBillSearch true "分页获取ReconciledBill列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rb/getReconciledBillList [get]
func (rbApi *ReconciledBillApi) GetReconciledBillList(c *gin.Context) {
	var pageInfo rebReq.ReconciledBillSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := rbService.GetReconciledBillInfoList(pageInfo); err != nil {
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

package bill

import (
	"server/global"
	"server/model/bill"
	"server/model/bill/request"
	cr "server/model/common/request"
	"server/model/common/response"

	"github.com/gin-gonic/gin"
)

type BillApi struct {
}

func (b *BillApi) CreateBill(c *gin.Context) {
	var r request.CreateBill
	if err := c.ShouldBindJSON(&r); err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := billService.CreateBill(r.Bill); err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(r, c)
}

func (b *BillApi) GetBillList(c *gin.Context) {
	var p cr.PageInfo
	var list interface{}
	var total int64
	var err error
	if err := c.ShouldBindJSON(&p); err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
	}
	switch p.KeyType {
	case "applicantId":
		list, total, err = billService.GetBillByApplicantId(p)
	case "approverId":
		list, total, err = billService.GetBillByApproverId(p)
	case "adminId":
		list, total, err = billService.GetBillByAdminId(p)
	case "initiatorID":
		list, total, err = billService.GetBillByInitiatorId(p)
	case "name":
		list, total, err = billService.GetBillListByName(p)
	default:
		list, total, err = billService.GetBillList(p)
	}
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     p.Page,
		PageSize: p.PageSize,
	}, "获取成功", c)
}
func (b *BillApi) UpdateBill(c *gin.Context) {
	var r bill.Bill
	if err := c.ShouldBindJSON(&r); err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := billService.UpdateBill(r); err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(r, c)
}
func (b *BillApi) DeleteBill(c *gin.Context) {
	var r request.DeleteBill
	if err := c.ShouldBindJSON(&r); err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := billService.DeleteBill(r.Bill); err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(r, c)
}

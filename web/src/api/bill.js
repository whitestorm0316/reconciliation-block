import service from '@/utils/request'

// @Tags Bill
// @Summary 分页获取角色列表
// @Security BillKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取用户列表"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bill/getBillList [post]
// {
//  page     int
//	pageSize int
// }
export const getBillList = (data) => {
  return service({
    url: '/bill/getBillList',
    method: 'post',
    data
  })
}

// @Tags Bill
// @Summary 创建基础Bill
// @Security BillKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Bill.CreateBillParams true "创建Bill"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bill/createBill [post]
export const createBill = (data) => {
  return service({
    url: '/bill/createBill',
    method: 'post',
    data
  })
}


// @Tags Bill
// @Summary 更新Bill
// @Security BillKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Bill.CreateBillParams true "更新Bill"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bill/updateBill [post]
export const updateBill = (data) => {
  return service({
    url: '/bill/updateBill',
    method: 'post',
    data
  })
}





// @Tags Bill
// @Summary 删除指定Bill
// @Security BillKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dbModel.Bill true "删除Bill"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bill/deleteBill [post]
export const deleteBill = (data) => {
  return service({
    url: '/bill/deleteBill',
    method: 'post',
    data
  })
}

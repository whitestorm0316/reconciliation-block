import service from '@/utils/request'

// @Tags ErrorBill
// @Summary 创建ErrorBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ErrorBill true "创建ErrorBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /eb/createErrorBill [post]
export const createErrorBill = (data) => {
  return service({
    url: '/eb/createErrorBill',
    method: 'post',
    data
  })
}

// @Tags ErrorBill
// @Summary 删除ErrorBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ErrorBill true "删除ErrorBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /eb/deleteErrorBill [delete]
export const deleteErrorBill = (data) => {
  return service({
    url: '/eb/deleteErrorBill',
    method: 'delete',
    data
  })
}

// @Tags ErrorBill
// @Summary 删除ErrorBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ErrorBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /eb/deleteErrorBill [delete]
export const deleteErrorBillByIds = (data) => {
  return service({
    url: '/eb/deleteErrorBillByIds',
    method: 'delete',
    data
  })
}

// @Tags ErrorBill
// @Summary 更新ErrorBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ErrorBill true "更新ErrorBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /eb/updateErrorBill [put]
export const updateErrorBill = (data) => {
  return service({
    url: '/eb/updateErrorBill',
    method: 'put',
    data
  })
}

// @Tags ErrorBill
// @Summary 用id查询ErrorBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ErrorBill true "用id查询ErrorBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /eb/findErrorBill [get]
export const findErrorBill = (params) => {
  return service({
    url: '/eb/findErrorBill',
    method: 'get',
    params
  })
}

// @Tags ErrorBill
// @Summary 分页获取ErrorBill列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ErrorBill列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /eb/getErrorBillList [get]
export const getErrorBillList = (params) => {
  return service({
    url: '/eb/getErrorBillList',
    method: 'get',
    params
  })
}

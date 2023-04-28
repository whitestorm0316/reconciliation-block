import service from '@/utils/request'

// @Tags LocalBill
// @Summary 创建LocalBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LocalBill true "创建LocalBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lb/createLocalBill [post]
export const createLocalBill = (data) => {
  return service({
    url: '/lb/createLocalBill',
    method: 'post',
    data
  })
}

// @Tags LocalBill
// @Summary 删除LocalBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LocalBill true "删除LocalBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lb/deleteLocalBill [delete]
export const deleteLocalBill = (data) => {
  return service({
    url: '/lb/deleteLocalBill',
    method: 'delete',
    data
  })
}

// @Tags LocalBill
// @Summary 删除LocalBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LocalBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lb/deleteLocalBill [delete]
export const deleteLocalBillByIds = (data) => {
  return service({
    url: '/lb/deleteLocalBillByIds',
    method: 'delete',
    data
  })
}

// @Tags LocalBill
// @Summary 更新LocalBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LocalBill true "更新LocalBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lb/updateLocalBill [put]
export const updateLocalBill = (data) => {
  return service({
    url: '/lb/updateLocalBill',
    method: 'put',
    data
  })
}

// @Tags LocalBill
// @Summary 用id查询LocalBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.LocalBill true "用id查询LocalBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lb/findLocalBill [get]
export const findLocalBill = (params) => {
  return service({
    url: '/lb/findLocalBill',
    method: 'get',
    params
  })
}

// @Tags LocalBill
// @Summary 分页获取LocalBill列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取LocalBill列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lb/getLocalBillList [get]
export const getLocalBillList = (params) => {
  return service({
    url: '/lb/getLocalBillList',
    method: 'get',
    params
  })
}

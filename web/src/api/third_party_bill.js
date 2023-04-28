import service from '@/utils/request'

// @Tags ThirdPartyBill
// @Summary 创建ThirdPartyBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ThirdPartyBill true "创建ThirdPartyBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tpb/createThirdPartyBill [post]
export const createThirdPartyBill = (data) => {
  return service({
    url: '/tpb/createThirdPartyBill',
    method: 'post',
    data
  })
}

// @Tags ThirdPartyBill
// @Summary 删除ThirdPartyBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ThirdPartyBill true "删除ThirdPartyBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tpb/deleteThirdPartyBill [delete]
export const deleteThirdPartyBill = (data) => {
  return service({
    url: '/tpb/deleteThirdPartyBill',
    method: 'delete',
    data
  })
}

// @Tags ThirdPartyBill
// @Summary 删除ThirdPartyBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ThirdPartyBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tpb/deleteThirdPartyBill [delete]
export const deleteThirdPartyBillByIds = (data) => {
  return service({
    url: '/tpb/deleteThirdPartyBillByIds',
    method: 'delete',
    data
  })
}

// @Tags ThirdPartyBill
// @Summary 更新ThirdPartyBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ThirdPartyBill true "更新ThirdPartyBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tpb/updateThirdPartyBill [put]
export const updateThirdPartyBill = (data) => {
  return service({
    url: '/tpb/updateThirdPartyBill',
    method: 'put',
    data
  })
}

// @Tags ThirdPartyBill
// @Summary 用id查询ThirdPartyBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ThirdPartyBill true "用id查询ThirdPartyBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tpb/findThirdPartyBill [get]
export const findThirdPartyBill = (params) => {
  return service({
    url: '/tpb/findThirdPartyBill',
    method: 'get',
    params
  })
}

// @Tags ThirdPartyBill
// @Summary 分页获取ThirdPartyBill列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ThirdPartyBill列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tpb/getThirdPartyBillList [get]
export const getThirdPartyBillList = (params) => {
  return service({
    url: '/tpb/getThirdPartyBillList',
    method: 'get',
    params
  })
}

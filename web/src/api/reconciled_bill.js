import service from '@/utils/request'

// @Tags ReconciledBill
// @Summary 创建ReconciledBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ReconciledBill true "创建ReconciledBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rb/createReconciledBill [post]
export const createReconciledBill = (data) => {
  return service({
    url: '/rb/createReconciledBill',
    method: 'post',
    data
  })
}

// @Tags ReconciledBill
// @Summary 删除ReconciledBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ReconciledBill true "删除ReconciledBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rb/deleteReconciledBill [delete]
export const deleteReconciledBill = (data) => {
  return service({
    url: '/rb/deleteReconciledBill',
    method: 'delete',
    data
  })
}

// @Tags ReconciledBill
// @Summary 删除ReconciledBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ReconciledBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rb/deleteReconciledBill [delete]
export const deleteReconciledBillByIds = (data) => {
  return service({
    url: '/rb/deleteReconciledBillByIds',
    method: 'delete',
    data
  })
}

// @Tags ReconciledBill
// @Summary 更新ReconciledBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ReconciledBill true "更新ReconciledBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /rb/updateReconciledBill [put]
export const updateReconciledBill = (data) => {
  return service({
    url: '/rb/updateReconciledBill',
    method: 'put',
    data
  })
}

// @Tags ReconciledBill
// @Summary 用id查询ReconciledBill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ReconciledBill true "用id查询ReconciledBill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /rb/findReconciledBill [get]
export const findReconciledBill = (params) => {
  return service({
    url: '/rb/findReconciledBill',
    method: 'get',
    params
  })
}

// @Tags ReconciledBill
// @Summary 分页获取ReconciledBill列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ReconciledBill列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rb/getReconciledBillList [get]
export const getReconciledBillList = (params) => {
  return service({
    url: '/rb/getReconciledBillList',
    method: 'get',
    params
  })
}


// @Tags SyncBlock
// @Summary 同步至区块链SyncBlock
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body reb.ReconciledBill true "同步至区块链SyncBlock"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /rb/syncBlock [post]

export const syncBlock = (data) => {
  return service({
    url: '/rb/syncBlock',
    method: 'post',
    data
  })
}

import service from '@/utils/request'

// @Tags Blockchain
// @Summary 创建Blockchain
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Blockchain true "创建Blockchain"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bc/createBlockchain [post]
export const createBlockchain = (data) => {
  return service({
    url: '/bc/createBlockchain',
    method: 'post',
    data
  })
}

// @Tags Blockchain
// @Summary 删除Blockchain
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Blockchain true "删除Blockchain"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bc/deleteBlockchain [delete]
export const deleteBlockchain = (data) => {
  return service({
    url: '/bc/deleteBlockchain',
    method: 'delete',
    data
  })
}

// @Tags Blockchain
// @Summary 删除Blockchain
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Blockchain"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /bc/deleteBlockchain [delete]
export const deleteBlockchainByIds = (data) => {
  return service({
    url: '/bc/deleteBlockchainByIds',
    method: 'delete',
    data
  })
}

// @Tags Blockchain
// @Summary 更新Blockchain
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Blockchain true "更新Blockchain"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /bc/updateBlockchain [put]
export const updateBlockchain = (data) => {
  return service({
    url: '/bc/updateBlockchain',
    method: 'put',
    data
  })
}

// @Tags Blockchain
// @Summary 用id查询Blockchain
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Blockchain true "用id查询Blockchain"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /bc/findBlockchain [get]
export const findBlockchain = (params) => {
  return service({
    url: '/bc/findBlockchain',
    method: 'get',
    params
  })
}

// @Tags Blockchain
// @Summary 分页获取Blockchain列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Blockchain列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /bc/getBlockchainList [get]
export const getBlockchainList = (params) => {
  return service({
    url: '/bc/getBlockchainList',
    method: 'get',
    params
  })
}

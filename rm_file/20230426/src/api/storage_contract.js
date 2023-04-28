import service from '@/utils/request'

// @Tags StorageContract
// @Summary 创建StorageContract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.StorageContract true "创建StorageContract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sc/createStorageContract [post]
export const createStorageContract = (data) => {
  return service({
    url: '/sc/createStorageContract',
    method: 'post',
    data
  })
}

// @Tags StorageContract
// @Summary 删除StorageContract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.StorageContract true "删除StorageContract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sc/deleteStorageContract [delete]
export const deleteStorageContract = (data) => {
  return service({
    url: '/sc/deleteStorageContract',
    method: 'delete',
    data
  })
}

// @Tags StorageContract
// @Summary 删除StorageContract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除StorageContract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sc/deleteStorageContract [delete]
export const deleteStorageContractByIds = (data) => {
  return service({
    url: '/sc/deleteStorageContractByIds',
    method: 'delete',
    data
  })
}

// @Tags StorageContract
// @Summary 更新StorageContract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.StorageContract true "更新StorageContract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sc/updateStorageContract [put]
export const updateStorageContract = (data) => {
  return service({
    url: '/sc/updateStorageContract',
    method: 'put',
    data
  })
}

// @Tags StorageContract
// @Summary 用id查询StorageContract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.StorageContract true "用id查询StorageContract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sc/findStorageContract [get]
export const findStorageContract = (params) => {
  return service({
    url: '/sc/findStorageContract',
    method: 'get',
    params
  })
}

// @Tags StorageContract
// @Summary 分页获取StorageContract列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取StorageContract列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sc/getStorageContractList [get]
export const getStorageContractList = (params) => {
  return service({
    url: '/sc/getStorageContractList',
    method: 'get',
    params
  })
}

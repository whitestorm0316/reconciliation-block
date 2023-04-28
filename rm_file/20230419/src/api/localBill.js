import service from '@/utils/request'

// @Tags Localbill
// @Summary 创建Localbill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Localbill true "创建Localbill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lb/createLocalbill [post]
export const createLocalbill = (data) => {
  return service({
    url: '/lb/createLocalbill',
    method: 'post',
    data
  })
}

// @Tags Localbill
// @Summary 删除Localbill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Localbill true "删除Localbill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lb/deleteLocalbill [delete]
export const deleteLocalbill = (data) => {
  return service({
    url: '/lb/deleteLocalbill',
    method: 'delete',
    data
  })
}

// @Tags Localbill
// @Summary 删除Localbill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Localbill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /lb/deleteLocalbill [delete]
export const deleteLocalbillByIds = (data) => {
  return service({
    url: '/lb/deleteLocalbillByIds',
    method: 'delete',
    data
  })
}

// @Tags Localbill
// @Summary 更新Localbill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Localbill true "更新Localbill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /lb/updateLocalbill [put]
export const updateLocalbill = (data) => {
  return service({
    url: '/lb/updateLocalbill',
    method: 'put',
    data
  })
}

// @Tags Localbill
// @Summary 用id查询Localbill
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Localbill true "用id查询Localbill"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /lb/findLocalbill [get]
export const findLocalbill = (params) => {
  return service({
    url: '/lb/findLocalbill',
    method: 'get',
    params
  })
}

// @Tags Localbill
// @Summary 分页获取Localbill列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Localbill列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /lb/getLocalbillList [get]
export const getLocalbillList = (params) => {
  return service({
    url: '/lb/getLocalbillList',
    method: 'get',
    params
  })
}

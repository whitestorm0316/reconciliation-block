package reb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"server/global"
	"server/model/common/request"
	"server/model/common/response"
	"server/model/reb"
	rebReq "server/model/reb/request"
	"server/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type StorageContractApi struct {
}

var scService = service.ServiceGroupApp.RebServiceGroup.StorageContractService

// CreateStorageContract 创建StorageContract
// @Tags StorageContract
// @Summary 创建StorageContract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body reb.StorageContract true "创建StorageContract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sc/createStorageContract [post]
func (scApi *StorageContractApi) CreateStorageContract(c *gin.Context) {

	type contract struct {
		global.GVA_MODEL
		Name       string `json:"name" form:"name" gorm:"column:name;comment:;"`
		Note       string `json:"note" form:"note" gorm:"column:note;comment:备注;"`
		Address    string `json:"address" form:"address" gorm:"column:address;comment:地址;"`
		BlockName  string `json:"blockName" form:"blockName" gorm:"column:block_name;comment:;"`
		Url        string `json:"url" form:"url" `
		PrivateKey string `json:"privateKey" form:"privateKey"`
	}
	var ct contract
	err := c.ShouldBindJSON(&ct)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	bs, _ := json.Marshal(ct)
	resp, err := http.Post("http://localhost:3000/contract/create", "application/json", bytes.NewReader(bs))

	if err != nil {
		fmt.Println(err)
	} else {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}

		global.GVA_LOG.Info(string(body))
		addr := string(body)
		sc := reb.StorageContract{
			Name:      ct.Name,
			Note:      ct.Note,
			Address:   addr,
			BlockName: ct.BlockName,
		}
		b, _ := json.Marshal(sc)
		global.GVA_LOG.Info(string(b))

		if err := scService.CreateStorageContract(sc); err != nil {
			global.GVA_LOG.Error("创建失败!", zap.Error(err))
			response.FailWithMessage("创建失败", c)
		} else {
			response.OkWithMessage("创建成功", c)
		}
	}

}

// DeleteStorageContract 删除StorageContract
// @Tags StorageContract
// @Summary 删除StorageContract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body reb.StorageContract true "删除StorageContract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sc/deleteStorageContract [delete]
func (scApi *StorageContractApi) DeleteStorageContract(c *gin.Context) {
	var sc reb.StorageContract
	err := c.ShouldBindJSON(&sc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := scService.DeleteStorageContract(sc); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteStorageContractByIds 批量删除StorageContract
// @Tags StorageContract
// @Summary 批量删除StorageContract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除StorageContract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sc/deleteStorageContractByIds [delete]
func (scApi *StorageContractApi) DeleteStorageContractByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := scService.DeleteStorageContractByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateStorageContract 更新StorageContract
// @Tags StorageContract
// @Summary 更新StorageContract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body reb.StorageContract true "更新StorageContract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sc/updateStorageContract [put]
func (scApi *StorageContractApi) UpdateStorageContract(c *gin.Context) {
	var sc reb.StorageContract
	err := c.ShouldBindJSON(&sc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := scService.UpdateStorageContract(sc); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindStorageContract 用id查询StorageContract
// @Tags StorageContract
// @Summary 用id查询StorageContract
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query reb.StorageContract true "用id查询StorageContract"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sc/findStorageContract [get]
func (scApi *StorageContractApi) FindStorageContract(c *gin.Context) {
	var sc reb.StorageContract
	err := c.ShouldBindQuery(&sc)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if resc, err := scService.GetStorageContract(sc.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resc": resc}, c)
	}
}

// GetStorageContractList 分页获取StorageContract列表
// @Tags StorageContract
// @Summary 分页获取StorageContract列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query rebReq.StorageContractSearch true "分页获取StorageContract列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sc/getStorageContractList [get]
func (scApi *StorageContractApi) GetStorageContractList(c *gin.Context) {
	var pageInfo rebReq.StorageContractSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := scService.GetStorageContractInfoList(pageInfo); err != nil {
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

package reb

import (
	"server/global"
	"server/model/reb"
	"server/model/common/request"
    rebReq "server/model/reb/request"
)

type StorageContractService struct {
}

// CreateStorageContract 创建StorageContract记录
// Author [piexlmax](https://github.com/piexlmax)
func (scService *StorageContractService) CreateStorageContract(sc reb.StorageContract) (err error) {
	err = global.GVA_DB.Create(&sc).Error
	return err
}

// DeleteStorageContract 删除StorageContract记录
// Author [piexlmax](https://github.com/piexlmax)
func (scService *StorageContractService)DeleteStorageContract(sc reb.StorageContract) (err error) {
	err = global.GVA_DB.Delete(&sc).Error
	return err
}

// DeleteStorageContractByIds 批量删除StorageContract记录
// Author [piexlmax](https://github.com/piexlmax)
func (scService *StorageContractService)DeleteStorageContractByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]reb.StorageContract{},"id in ?",ids.Ids).Error
	return err
}

// UpdateStorageContract 更新StorageContract记录
// Author [piexlmax](https://github.com/piexlmax)
func (scService *StorageContractService)UpdateStorageContract(sc reb.StorageContract) (err error) {
	err = global.GVA_DB.Save(&sc).Error
	return err
}

// GetStorageContract 根据id获取StorageContract记录
// Author [piexlmax](https://github.com/piexlmax)
func (scService *StorageContractService)GetStorageContract(id uint) (sc reb.StorageContract, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&sc).Error
	return
}

// GetStorageContractInfoList 分页获取StorageContract记录
// Author [piexlmax](https://github.com/piexlmax)
func (scService *StorageContractService)GetStorageContractInfoList(info rebReq.StorageContractSearch) (list []reb.StorageContract, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&reb.StorageContract{})
    var scs []reb.StorageContract
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.Note != "" {
        db = db.Where("note LIKE ?","%"+ info.Note+"%")
    }
    if info.Address != "" {
        db = db.Where("address LIKE ?","%"+ info.Address+"%")
    }
    if info.PrivateKey != "" {
        db = db.Where("private_key LIKE ?","%"+ info.PrivateKey+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&scs).Error
	return  scs, total, err
}

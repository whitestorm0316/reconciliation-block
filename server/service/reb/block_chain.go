package reb

import (
	"server/global"
	"server/model/reb"
	"server/model/common/request"
    rebReq "server/model/reb/request"
)

type BlockchainService struct {
}

// CreateBlockchain 创建Blockchain记录
// Author [piexlmax](https://github.com/piexlmax)
func (bcService *BlockchainService) CreateBlockchain(bc reb.Blockchain) (err error) {
	err = global.GVA_DB.Create(&bc).Error
	return err
}

// DeleteBlockchain 删除Blockchain记录
// Author [piexlmax](https://github.com/piexlmax)
func (bcService *BlockchainService)DeleteBlockchain(bc reb.Blockchain) (err error) {
	err = global.GVA_DB.Delete(&bc).Error
	return err
}

// DeleteBlockchainByIds 批量删除Blockchain记录
// Author [piexlmax](https://github.com/piexlmax)
func (bcService *BlockchainService)DeleteBlockchainByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]reb.Blockchain{},"id in ?",ids.Ids).Error
	return err
}

// UpdateBlockchain 更新Blockchain记录
// Author [piexlmax](https://github.com/piexlmax)
func (bcService *BlockchainService)UpdateBlockchain(bc reb.Blockchain) (err error) {
	err = global.GVA_DB.Save(&bc).Error
	return err
}

// GetBlockchain 根据id获取Blockchain记录
// Author [piexlmax](https://github.com/piexlmax)
func (bcService *BlockchainService)GetBlockchain(id uint) (bc reb.Blockchain, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&bc).Error
	return
}

// GetBlockchainInfoList 分页获取Blockchain记录
// Author [piexlmax](https://github.com/piexlmax)
func (bcService *BlockchainService)GetBlockchainInfoList(info rebReq.BlockchainSearch) (list []reb.Blockchain, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&reb.Blockchain{})
    var bcs []reb.Blockchain
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
    if info.Url != "" {
        db = db.Where("url LIKE ?","%"+ info.Url+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&bcs).Error
	return  bcs, total, err
}

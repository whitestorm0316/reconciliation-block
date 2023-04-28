package test

import (
	"server/global"
	"server/model/test"
	"server/model/common/request"
    testReq "server/model/test/request"
)

type LocalbillService struct {
}

// CreateLocalbill 创建Localbill记录
// Author [piexlmax](https://github.com/piexlmax)
func (lbService *LocalbillService) CreateLocalbill(lb test.Localbill) (err error) {
	err = global.GVA_DB.Create(&lb).Error
	return err
}

// DeleteLocalbill 删除Localbill记录
// Author [piexlmax](https://github.com/piexlmax)
func (lbService *LocalbillService)DeleteLocalbill(lb test.Localbill) (err error) {
	err = global.GVA_DB.Delete(&lb).Error
	return err
}

// DeleteLocalbillByIds 批量删除Localbill记录
// Author [piexlmax](https://github.com/piexlmax)
func (lbService *LocalbillService)DeleteLocalbillByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]test.Localbill{},"id in ?",ids.Ids).Error
	return err
}

// UpdateLocalbill 更新Localbill记录
// Author [piexlmax](https://github.com/piexlmax)
func (lbService *LocalbillService)UpdateLocalbill(lb test.Localbill) (err error) {
	err = global.GVA_DB.Save(&lb).Error
	return err
}

// GetLocalbill 根据id获取Localbill记录
// Author [piexlmax](https://github.com/piexlmax)
func (lbService *LocalbillService)GetLocalbill(id uint) (lb test.Localbill, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&lb).Error
	return
}

// GetLocalbillInfoList 分页获取Localbill记录
// Author [piexlmax](https://github.com/piexlmax)
func (lbService *LocalbillService)GetLocalbillInfoList(info testReq.LocalbillSearch) (list []test.Localbill, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&test.Localbill{})
    var lbs []test.Localbill
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.TradingUserId != nil {
        db = db.Where("trading_user_id = ?",info.TradingUserId)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&lbs).Error
	return  lbs, total, err
}

package reb

import (
	"server/global"
	"server/model/reb"
	"server/model/common/request"
    rebReq "server/model/reb/request"
)

type ThirdPartyBillService struct {
}

// CreateThirdPartyBill 创建ThirdPartyBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (tpbService *ThirdPartyBillService) CreateThirdPartyBill(tpb reb.ThirdPartyBill) (err error) {
	err = global.GVA_DB.Create(&tpb).Error
	return err
}

// DeleteThirdPartyBill 删除ThirdPartyBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (tpbService *ThirdPartyBillService)DeleteThirdPartyBill(tpb reb.ThirdPartyBill) (err error) {
	err = global.GVA_DB.Delete(&tpb).Error
	return err
}

// DeleteThirdPartyBillByIds 批量删除ThirdPartyBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (tpbService *ThirdPartyBillService)DeleteThirdPartyBillByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]reb.ThirdPartyBill{},"id in ?",ids.Ids).Error
	return err
}

// UpdateThirdPartyBill 更新ThirdPartyBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (tpbService *ThirdPartyBillService)UpdateThirdPartyBill(tpb reb.ThirdPartyBill) (err error) {
	err = global.GVA_DB.Save(&tpb).Error
	return err
}

// GetThirdPartyBill 根据id获取ThirdPartyBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (tpbService *ThirdPartyBillService)GetThirdPartyBill(id uint) (tpb reb.ThirdPartyBill, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&tpb).Error
	return
}

// GetThirdPartyBillInfoList 分页获取ThirdPartyBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (tpbService *ThirdPartyBillService)GetThirdPartyBillInfoList(info rebReq.ThirdPartyBillSearch) (list []reb.ThirdPartyBill, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&reb.ThirdPartyBill{})
    var tpbs []reb.ThirdPartyBill
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.TransactionNumber != "" {
        db = db.Where("transaction_number LIKE ?","%"+ info.TransactionNumber+"%")
    }
        if info.StartTransactionTime != nil && info.EndTransactionTime != nil {
            db = db.Where("transaction_time BETWEEN ? AND ? ",info.StartTransactionTime,info.EndTransactionTime)
        }
    if info.TransactionType != "" {
        db = db.Where("transaction_type LIKE ?","%"+ info.TransactionType+"%")
    }
    if info.Counterparty != "" {
        db = db.Where("counterparty LIKE ?","%"+ info.Counterparty+"%")
    }
    if info.PaymentMethod != "" {
        db = db.Where("payment_method LIKE ?","%"+ info.PaymentMethod+"%")
    }
    if info.IncomeOrExpenses != "" {
        db = db.Where("income_or_expenses LIKE ?","%"+ info.IncomeOrExpenses+"%")
    }
        if info.StartMoney != nil && info.EndMoney != nil {
            db = db.Where("money BETWEEN ? AND ? ",info.StartMoney,info.EndMoney)
        }
    if info.Reconciled != nil {
        db = db.Where("reconciled = ?",info.Reconciled)
    }
    if info.Organization != "" {
        db = db.Where("organization LIKE ?","%"+ info.Organization+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&tpbs).Error
	return  tpbs, total, err
}

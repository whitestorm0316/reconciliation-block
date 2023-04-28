package reb

import (
	"server/global"
	"server/model/reb"
	"server/model/common/request"
    rebReq "server/model/reb/request"
)

type LocalBillService struct {
}

// CreateLocalBill 创建LocalBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (lbService *LocalBillService) CreateLocalBill(lb reb.LocalBill) (err error) {
	err = global.GVA_DB.Create(&lb).Error
	return err
}

// DeleteLocalBill 删除LocalBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (lbService *LocalBillService)DeleteLocalBill(lb reb.LocalBill) (err error) {
	err = global.GVA_DB.Delete(&lb).Error
	return err
}

// DeleteLocalBillByIds 批量删除LocalBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (lbService *LocalBillService)DeleteLocalBillByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]reb.LocalBill{},"id in ?",ids.Ids).Error
	return err
}

// UpdateLocalBill 更新LocalBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (lbService *LocalBillService)UpdateLocalBill(lb reb.LocalBill) (err error) {
	err = global.GVA_DB.Save(&lb).Error
	return err
}

// GetLocalBill 根据id获取LocalBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (lbService *LocalBillService)GetLocalBill(id uint) (lb reb.LocalBill, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&lb).Error
	return
}

// GetLocalBillInfoList 分页获取LocalBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (lbService *LocalBillService)GetLocalBillInfoList(info rebReq.LocalBillSearch) (list []reb.LocalBill, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&reb.LocalBill{})
    var lbs []reb.LocalBill
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.TransactionUserId != nil {
        db = db.Where("trading_user_id = ?",info.TransactionUserId)
    }
    if info.TransactionNumber != "" {
        db = db.Where("transaction_number = ?",info.TransactionNumber)
    }
        if info.StartTransactionTime != nil && info.EndTransactionTime != nil {
            db = db.Where("transaction_time BETWEEN ? AND ? ",info.StartTransactionTime,info.EndTransactionTime)
        }
    if info.TransactionType != "" {
        db = db.Where("transaction_type = ?",info.TransactionType)
    }
    if info.Counterparty != "" {
        db = db.Where("counterparty = ?",info.Counterparty)
    }
    if info.PaymentMethod != "" {
        db = db.Where("payment_method = ?",info.PaymentMethod)
    }
    if info.IncomeOrExpenses != "" {
        db = db.Where("income_or_expenses = ?",info.IncomeOrExpenses)
    }
        if info.StartMoney != nil && info.EndMoney != nil {
            db = db.Where("money BETWEEN ? AND ? ",info.StartMoney,info.EndMoney)
        }
    if info.Reconciled != nil {
        db = db.Where("reconciled = ?",info.Reconciled)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&lbs).Error
	return  lbs, total, err
}

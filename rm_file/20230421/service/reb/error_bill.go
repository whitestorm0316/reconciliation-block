package reb

import (
	"server/global"
	"server/model/common/request"
	"server/model/reb"
	rebReq "server/model/reb/request"
)

type ErrorBillService struct {
}

// CreateErrorBill 创建ErrorBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (ebService *ErrorBillService) CreateErrorBill(eb reb.ErrorBill) (err error) {
	err = global.GVA_DB.Create(&eb).Error
	return err
}

// DeleteErrorBill 删除ErrorBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (ebService *ErrorBillService) DeleteErrorBill(eb reb.ErrorBill) (err error) {
	err = global.GVA_DB.Delete(&eb).Error
	return err
}

// DeleteErrorBillByIds 批量删除ErrorBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (ebService *ErrorBillService) DeleteErrorBillByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]reb.ErrorBill{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateErrorBill 更新ErrorBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (ebService *ErrorBillService) UpdateErrorBill(eb reb.ErrorBill) (err error) {
	err = global.GVA_DB.Save(&eb).Error
	return err
}

// GetErrorBill 根据id获取ErrorBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (ebService *ErrorBillService) GetErrorBill(id uint) (eb reb.ErrorBill, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&eb).Error
	return
}

// GetErrorBillInfoList 分页获取ErrorBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (ebService *ErrorBillService) GetErrorBillInfoList(info rebReq.ErrorBillSearch) (list []reb.ErrorBill, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&reb.ErrorBill{})
	var ebs []reb.ErrorBill
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.TransactionUserId != nil {
		db = db.Where("trading_user_id = ?", info.TransactionUserId)
	}
	if info.TransactionNumber != "" {
		db = db.Where("transaction_number LIKE ?", "%"+info.TransactionNumber+"%")
	}
	if info.StartTransactionTime != nil && info.EndTransactionTime != nil {
		db = db.Where("transaction_time BETWEEN ? AND ? ", info.StartTransactionTime, info.EndTransactionTime)
	}
	if info.TransactionType != "" {
		db = db.Where("transaction_type LIKE ?", "%"+info.TransactionType+"%")
	}
	if info.Counterparty != "" {
		db = db.Where("counterparty LIKE ?", "%"+info.Counterparty+"%")
	}
	if info.PaymentMethod != "" {
		db = db.Where("payment_method LIKE ?", "%"+info.PaymentMethod+"%")
	}
	if info.IncomeOrExpenses != "" {
		db = db.Where("income_or_expenses LIKE ?", "%"+info.IncomeOrExpenses+"%")
	}
	if info.StartMoney != nil && info.EndMoney != nil {
		db = db.Where("money BETWEEN ? AND ? ", info.StartMoney, info.EndMoney)
	}
	if info.Reconciled != nil {
		db = db.Where("reconciled = ?", info.Reconciled)
	}
	if info.Organization != "" {
		db = db.Where("organization LIKE ?", "%"+info.Organization+"%")
	}
	if info.Reconciler != nil {
		db = db.Where("reconciler = ?", info.Reconciler)
	}
	if info.ReviewerId != nil {
		db = db.Where("reviewer = ?", info.ReviewerId)
	}
	if info.Progress != "" {
		db = db.Where("progress = ?", info.Progress)
	}
	if info.Reason != "" {
		db = db.Where("reason LIKE ?", "%"+info.Reason+"%")
	}
	if info.ErrorReason != "" {
		db = db.Where("error_reason LIKE ?", "%"+info.ErrorReason+"%")
	}
	if info.Note != "" {
		db = db.Where("note LIKE ?", "%"+info.Note+"%")
	}
	if info.LocalBillId != nil {
		db = db.Where("local_bill_id = ?", info.LocalBillId)
	}
	if info.ThirdPartyBillId != nil {
		db = db.Where("third_party_bill_id = ?", info.ThirdPartyBillId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&ebs).Error
	return ebs, total, err
}

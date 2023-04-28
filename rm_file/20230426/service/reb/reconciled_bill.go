package reb

import (
	"server/global"
	"server/model/reb"
	"server/model/common/request"
    rebReq "server/model/reb/request"
)

type ReconciledBillService struct {
}

// CreateReconciledBill 创建ReconciledBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (rbService *ReconciledBillService) CreateReconciledBill(rb reb.ReconciledBill) (err error) {
	err = global.GVA_DB.Create(&rb).Error
	return err
}

// DeleteReconciledBill 删除ReconciledBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (rbService *ReconciledBillService)DeleteReconciledBill(rb reb.ReconciledBill) (err error) {
	err = global.GVA_DB.Delete(&rb).Error
	return err
}

// DeleteReconciledBillByIds 批量删除ReconciledBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (rbService *ReconciledBillService)DeleteReconciledBillByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]reb.ReconciledBill{},"id in ?",ids.Ids).Error
	return err
}

// UpdateReconciledBill 更新ReconciledBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (rbService *ReconciledBillService)UpdateReconciledBill(rb reb.ReconciledBill) (err error) {
	err = global.GVA_DB.Save(&rb).Error
	return err
}

// GetReconciledBill 根据id获取ReconciledBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (rbService *ReconciledBillService)GetReconciledBill(id uint) (rb reb.ReconciledBill, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&rb).Error
	return
}

// GetReconciledBillInfoList 分页获取ReconciledBill记录
// Author [piexlmax](https://github.com/piexlmax)
func (rbService *ReconciledBillService)GetReconciledBillInfoList(info rebReq.ReconciledBillSearch) (list []reb.ReconciledBill, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&reb.ReconciledBill{})
    var rbs []reb.ReconciledBill
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.TransactionUserId != nil {
        db = db.Where("trading_user_id = ?",info.TransactionUserId)
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
    if info.Organization != "" {
        db = db.Where("organization LIKE ?","%"+ info.Organization+"%")
    }
    if info.Reconciler != nil {
        db = db.Where("reconciler = ?",info.Reconciler)
    }
    if info.ReviewerId != nil {
        db = db.Where("reviewer = ?",info.ReviewerId)
    }
    if info.Note != "" {
        db = db.Where("note LIKE ?","%"+ info.Note+"%")
    }
    if info.LocalBillId != nil {
        db = db.Where("local_bill_id = ?",info.LocalBillId)
    }
    if info.ThirdPartyBillId != nil {
        db = db.Where("third_party_bill_id = ?",info.ThirdPartyBillId)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&rbs).Error
	return  rbs, total, err
}

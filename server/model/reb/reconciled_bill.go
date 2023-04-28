// 自动生成模板ReconciledBill
package reb

import (
	"server/global"
	"time"
)

// ReconciledBill 结构体
type ReconciledBill struct {
      global.GVA_MODEL
      TransactionUserId  *int `json:"tradingUserId" form:"tradingUserId" gorm:"column:trading_user_id;comment:交易员工ID;"`
      TransactionNumber  string `json:"transactionNumber" form:"transactionNumber" gorm:"column:transaction_number;comment:交易ID;"`
      TransactionTime  *time.Time `json:"transactionTime" form:"transactionTime" gorm:"column:transaction_time;comment:交易时间;"`
      TransactionType  string `json:"transactionType" form:"transactionType" gorm:"column:transaction_type;comment:交易类型;"`
      Counterparty  string `json:"counterparty" form:"counterparty" gorm:"column:counterparty;comment:交易对方;"`
      PaymentMethod  string `json:"paymentMethod" form:"paymentMethod" gorm:"column:payment_method;comment:;"`
      IncomeOrExpenses  string `json:"incomeOrExpenses" form:"incomeOrExpenses" gorm:"column:income_or_expenses;comment:收入或者支出;"`
      Money  *float64 `json:"money" form:"money" gorm:"column:money;comment:金额;"`
      Organization  string `json:"organization" form:"organization" gorm:"column:organization;comment:账单所属机构;"`
      Reconciler  *int `json:"reconciler" form:"reconciler" gorm:"column:reconciler;comment:对账员ID;"`
      ReviewerId  *int `json:"reviewer" form:"reviewer" gorm:"column:reviewer;comment:审核员ID;"`
      Note  string `json:"note" form:"note" gorm:"column:note;comment:备注;"`
      LocalBillId  *int `json:"localBillId" form:"localBillId" gorm:"column:local_bill_id;comment:本地账单ID;"`
      ThirdPartyBillId  *int `json:"thirdPartyBillId" form:"thirdPartyBillId" gorm:"column:third_party_bill_id;comment:;"`
      Blocked  *bool `json:"blocked" form:"blocked" gorm:"column:blocked;comment:是否上传至区块链;"`
      ContractName  string `json:"contractName" form:"contractName" gorm:"column:contract_name;comment:合约名称;"`
}


// TableName ReconciledBill 表名
func (ReconciledBill) TableName() string {
  return "reconciled_bill"
}


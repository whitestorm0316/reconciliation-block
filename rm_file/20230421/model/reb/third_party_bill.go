// 自动生成模板ThirdPartyBill
package reb

import (
	"server/global"
	"time"
)

// ThirdPartyBill 结构体
type ThirdPartyBill struct {
      global.GVA_MODEL
      TransactionNumber  string `json:"transactionNumber" form:"transactionNumber" gorm:"column:transaction_number;comment:交易ID;"`
      TransactionTime  *time.Time `json:"transactionTime" form:"transactionTime" gorm:"column:transaction_time;comment:交易时间;"`
      TransactionType  string `json:"transactionType" form:"transactionType" gorm:"column:transaction_type;comment:交易类型;"`
      Counterparty  string `json:"counterparty" form:"counterparty" gorm:"column:counterparty;comment:交易对方;"`
      PaymentMethod  string `json:"paymentMethod" form:"paymentMethod" gorm:"column:payment_method;comment:;"`
      IncomeOrExpenses  string `json:"incomeOrExpenses" form:"incomeOrExpenses" gorm:"column:income_or_expenses;comment:收入或者支出;"`
      Money  *float64 `json:"money" form:"money" gorm:"column:money;comment:金额;"`
      Reconciled  *bool `json:"reconciled" form:"reconciled" gorm:"column:reconciled;comment:是否已对账;"`
      Organization  string `json:"organization" form:"organization" gorm:"column:organization;comment:账单所属机构;"`
}


// TableName ThirdPartyBill 表名
func (ThirdPartyBill) TableName() string {
  return "third_party_bill"
}


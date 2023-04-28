// 自动生成模板LocalBill
package reb

import (
	"server/global"
	"time"
)

// LocalBill 结构体
type LocalBill struct {
      global.GVA_MODEL
      TransactionUserId  *int `json:"tradingUserId" form:"tradingUserId" gorm:"column:trading_user_id;comment:交易员工ID;"`
      TransactionNumber  string `json:"transactionNumber" form:"transactionNumber" gorm:"column:transaction_number;comment:交易ID;"`
      TransactionTime  *time.Time `json:"transactionTime" form:"transactionTime" gorm:"column:transaction_time;comment:交易时间;"`
      TransactionType  string `json:"transactionType" form:"transactionType" gorm:"column:transaction_type;comment:交易类型;"`
      Counterparty  string `json:"counterparty" form:"counterparty" gorm:"column:counterparty;comment:交易对方;"`
      PaymentMethod  string `json:"paymentMethod" form:"paymentMethod" gorm:"column:payment_method;comment:;"`
      IncomeOrExpenses  string `json:"incomeOrExpenses" form:"incomeOrExpenses" gorm:"column:income_or_expenses;comment:收入或者支出;"`
      Money  *float64 `json:"money" form:"money" gorm:"column:money;comment:金额;"`
}


// TableName LocalBill 表名
func (LocalBill) TableName() string {
  return "local_bill"
}


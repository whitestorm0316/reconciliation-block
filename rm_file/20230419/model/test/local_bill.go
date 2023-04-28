// 自动生成模板Localbill
package test

import (
	"server/global"
	
)

// Localbill 结构体
type Localbill struct {
      global.GVA_MODEL
      TradingUserId  *int `json:"tradingUserId" form:"tradingUserId" gorm:"column:trading_user_id;comment:交易员工ID;"`
}


// TableName Localbill 表名
func (Localbill) TableName() string {
  return "localbill"
}


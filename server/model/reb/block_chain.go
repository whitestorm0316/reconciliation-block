// 自动生成模板Blockchain
package reb

import (
	"server/global"
	
)

// Blockchain 结构体
type Blockchain struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;"`
      Note  string `json:"note" form:"note" gorm:"column:note;comment:备注;"`
      Address  string `json:"address" form:"address" gorm:"column:address;comment:地址;"`
      PrivateKey  string `json:"privateKey" form:"privateKey" gorm:"column:private_key;comment:;"`
      Url  string `json:"url" form:"url" gorm:"column:url;comment:区块链链接地址;"`
}


// TableName Blockchain 表名
func (Blockchain) TableName() string {
  return "block_chain"
}


// 自动生成模板StorageContract
package reb

import (
	"server/global"
	
)

// StorageContract 结构体
type StorageContract struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;"`
      Note  string `json:"note" form:"note" gorm:"column:note;comment:备注;"`
      Address  string `json:"address" form:"address" gorm:"column:address;comment:地址;"`
      PrivateKey  string `json:"privateKey" form:"privateKey" gorm:"column:private_key;comment:;"`
}


// TableName StorageContract 表名
func (StorageContract) TableName() string {
  return "storage_contract"
}


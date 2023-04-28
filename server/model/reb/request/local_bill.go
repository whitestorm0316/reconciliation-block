package request

import (
	"server/model/reb"
	"server/model/common/request"
	"time"
)

type LocalBillSearch struct{
    reb.LocalBill
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    StartTransactionTime  *time.Time  `json:"startTransactionTime" form:"startTransactionTime"`
    EndTransactionTime  *time.Time  `json:"endTransactionTime" form:"endTransactionTime"`
    StartMoney  *float64  `json:"startMoney" form:"startMoney"`
    EndMoney  *float64  `json:"endMoney" form:"endMoney"`
    request.PageInfo
}

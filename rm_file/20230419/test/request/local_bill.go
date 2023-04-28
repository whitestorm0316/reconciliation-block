package request

import (
	"server/model/test"
	"server/model/common/request"
	"time"
)

type LocalbillSearch struct{
    test.Localbill
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}

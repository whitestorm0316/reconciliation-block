package request

import (
	"server/model/reb"
	"server/model/common/request"
	"time"
)

type BlockchainSearch struct{
    reb.Blockchain
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}

package v1

import (
	"server/api/v1/bill"
	"server/api/v1/example"
	"server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	BillApiGroup    bill.ApiGroup
}

var ApiGroupApp = new(ApiGroup)

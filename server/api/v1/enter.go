package v1

import (
	"server/api/v1/bill"
	"server/api/v1/example"
	"server/api/v1/reb"
	"server/api/v1/system"
	"server/api/v1/test"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	BillApiGroup    bill.ApiGroup
	TestApiGroup    test.ApiGroup
	RebApiGroup     reb.ApiGroup
}

var ApiGroupApp = new(ApiGroup)

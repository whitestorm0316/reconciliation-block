package service

import (
	"server/service/bill"
	"server/service/example"
	"server/service/reb"
	"server/service/system"
	"server/service/test"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	BillServiceGroup    bill.ServiceGroup
	TestServiceGroup    test.ServiceGroup
	RebServiceGroup     reb.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)

package service

import (
	"server/service/bill"
	"server/service/example"
	"server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	BillServiceGroup    bill.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)

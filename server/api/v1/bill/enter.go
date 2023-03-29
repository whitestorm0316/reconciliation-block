package bill

import "server/service"

type ApiGroup struct {
	BillApi
}

var (
	billService = service.ServiceGroupApp.BillServiceGroup.BillService
)

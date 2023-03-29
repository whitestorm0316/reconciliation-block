package router

import (
	"server/router/bill"
	"server/router/example"
	"server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Bill    bill.RouterGroup
}

var RouterGroupApp = new(RouterGroup)

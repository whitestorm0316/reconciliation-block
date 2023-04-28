package router

import (
	"server/router/bill"
	"server/router/example"
	"server/router/reb"
	"server/router/system"
	"server/router/test"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Bill    bill.RouterGroup
	Test    test.RouterGroup
	Reb     reb.RouterGroup
}

var RouterGroupApp = new(RouterGroup)

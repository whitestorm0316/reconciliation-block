package initialize

import (
	_ "server/model/bill"
	_ "server/source/example"
	_ "server/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}

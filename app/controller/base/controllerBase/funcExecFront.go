package controllerBase

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
)

func (self *EntityControllerBase) execFront() {
	routeType := self.Derived.GetRouteType()
	self.RouteType = routeType
	switch routeType {
	case conf.RouteTypeFile:
		self.execFile()
	case conf.RouteTypeApi:
		self.execApi()
	case conf.RouteTypeWeb:
		self.execWeb()
	default:
		msgError := fmt.Sprintf(
			"The routeType |%s| is invalid",
			routeType)
		err.ErrPanic(msgError)
	}
	return
}

func (self *EntityControllerBase) GetRouteType() (routeType int) {
	routeType = conf.RouteTypeDefault
	return routeType
}

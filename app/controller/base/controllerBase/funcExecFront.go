package controllerBase

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/consts"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
)

func (self *EntityControllerBase) execFront() {
	entityDataController := self.Derived.GetEntityDataController()
	routeType := entityDataController.GetRouteType()
	self.RouteType = routeType
	switch routeType {
	case consts.RouteTypeFile:
		self.execFile()
	case consts.RouteTypeApi:
		self.execApi()
	case consts.RouteTypeWeb:
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

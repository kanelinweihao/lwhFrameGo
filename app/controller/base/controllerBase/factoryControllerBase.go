package controllerBase

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
)

func (self *EntityControllerBase) Load(entityControllerDerived typeInterface.EntityController) typeInterface.EntityController {
	self.Derived = entityControllerDerived
	return self.Derived
}

func (self *EntityControllerBase) Init(routeName string) typeInterface.EntityController {
	self.RouteName = routeName
	return self.Derived
}

func (self *EntityControllerBase) Close() {
	if self.Req == nil {
		return
	}
	errClose := self.Req.Body.Close()
	err.ErrCheck(errClose)
	self.Req = nil
	self.Resp = nil
	return
}

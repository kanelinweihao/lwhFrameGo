package controllerBase

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

func (self *EntityControllerBase) Load(entityControllerDerived typeStruct.EntityController) typeStruct.EntityController {
	self.Derived = entityControllerDerived
	return self.Derived
}

func (self *EntityControllerBase) Init(routeName string) typeStruct.EntityController {
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
	self.ValuesFromReq = nil
	return
}

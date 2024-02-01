package controllerApiUserGetUserIdByMobileNoAndOrgId

import (
	"github.com/kanelinweihao/lwhFrameGo/app/controller/base/controllerBase"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
)

var routeName = "/api/user/getUserIdByMobileNoAndOrgId"

func InitEntityController() (entityController typeInterface.EntityController) {
	entityController = new(EntityController)
	entityControllerBase := new(controllerBase.EntityControllerBase)
	entityController.Load(entityControllerBase).Init(routeName)
	return entityController
}

func (self *EntityController) Load(entityControllerBase typeInterface.EntityController) typeInterface.EntityController {
	self.EntityController = entityControllerBase
	self.EntityController.Load(self)
	return self
}

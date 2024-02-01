package controllerApiAuthLogin

import (
	"github.com/kanelinweihao/lwhFrameGo/app/controller/base/controllerBase"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
)

var routeName = "/api/auth/login"

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
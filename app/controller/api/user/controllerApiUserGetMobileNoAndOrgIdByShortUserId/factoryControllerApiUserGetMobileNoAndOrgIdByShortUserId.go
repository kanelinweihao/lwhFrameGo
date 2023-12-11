package controllerApiUserGetMobileNoAndOrgIdByShortUserId

import (
	"github.com/kanelinweihao/lwhFrameGo/app/controller/base/controllerBase"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

var routeName = "/api/user/getMobileNoAndOrgIdByShortUserId"

func InitEntityController() (entityController typeStruct.EntityController) {
	entityController = new(EntityController)
	entityControllerBase := new(controllerBase.EntityControllerBase)
	entityController.Load(entityControllerBase).Init(routeName)
	return entityController
}

func (self *EntityController) Load(entityControllerBase typeStruct.EntityController) typeStruct.EntityController {
	self.EntityController = entityControllerBase
	self.EntityController.Load(self)
	return self
}

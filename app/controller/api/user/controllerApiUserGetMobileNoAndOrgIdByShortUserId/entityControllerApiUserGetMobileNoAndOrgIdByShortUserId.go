package controllerApiUserGetMobileNoAndOrgIdByShortUserId

import (
	"github.com/kanelinweihao/lwhFrameGo/app/controller/web/user/controllerWebUserGetMobileNoAndOrgIdByShortUserId"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

type EntityController struct {
	typeInterface.EntityController
}

func (self *EntityController) GetEntityDataController() (entityDataController *typeStruct.EntityDataController) {
	entityDataController = controllerWebUserGetMobileNoAndOrgIdByShortUserId.EntityDataController.ToEntityAPI()
	return entityDataController
}

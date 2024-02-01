package controllerApiProjectTitle

import (
	"github.com/kanelinweihao/lwhFrameGo/app/controller/web/common/controllerWebProjectTitle"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

type EntityController struct {
	typeInterface.EntityController
}

func (self *EntityController) GetEntityDataController() (entityDataController *typeStruct.EntityDataController) {
	entityDataController = controllerWebProjectTitle.EntityDataController.ToEntityAPI()
	return entityDataController
}

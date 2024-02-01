package controllerApiProjectVersion

import (
	"github.com/kanelinweihao/lwhFrameGo/app/controller/web/common/controllerWebProjectVersion"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

type EntityController struct {
	typeInterface.EntityController
}

func (self *EntityController) GetEntityDataController() (entityDataController *typeStruct.EntityDataController) {
	entityDataController = controllerWebProjectVersion.EntityDataController.ToEntityAPI()
	return entityDataController
}

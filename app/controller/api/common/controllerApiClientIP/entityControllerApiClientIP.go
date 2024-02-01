package controllerApiClientIP

import (
	"github.com/kanelinweihao/lwhFrameGo/app/controller/web/common/controllerWebClientIP"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

type EntityController struct {
	typeInterface.EntityController
}

func (self *EntityController) GetEntityDataController() (entityDataController *typeStruct.EntityDataController) {
	entityDataController = controllerWebClientIP.EntityDataController.ToEntityAPI()
	return entityDataController
}

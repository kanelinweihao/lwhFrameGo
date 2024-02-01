package controllerApiRateGetIRRByArrAmount

import (
	"github.com/kanelinweihao/lwhFrameGo/app/controller/web/rate/controllerWebRateGetIRRByArrAmount"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

type EntityController struct {
	typeInterface.EntityController
}

func (self *EntityController) GetEntityDataController() (entityDataController *typeStruct.EntityDataController) {
	entityDataController = controllerWebRateGetIRRByArrAmount.EntityDataController.ToEntityAPI()
	return entityDataController
}

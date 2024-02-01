package controllerApiCustomerBehaviorGetCustomerBehaviorByUserId

import (
	"github.com/kanelinweihao/lwhFrameGo/app/controller/web/customerBehavior/controllerWebCustomerBehaviorGetCustomerBehaviorByUserId"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

type EntityController struct {
	typeInterface.EntityController
}

func (self *EntityController) GetEntityDataController() (entityDataController *typeStruct.EntityDataController) {
	entityDataController = controllerWebCustomerBehaviorGetCustomerBehaviorByUserId.EntityDataController.ToEntityAPI()
	return entityDataController
}

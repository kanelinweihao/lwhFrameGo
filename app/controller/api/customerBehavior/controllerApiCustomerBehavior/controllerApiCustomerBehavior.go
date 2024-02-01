package controllerApiCustomerBehavior

import (
	"github.com/kanelinweihao/lwhFrameGo/app/controller/api/customerBehavior/controllerApiCustomerBehaviorGetCustomerBehaviorByUserId"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
	"net/http"
)

var entityController typeInterface.EntityController

func GetCustomerBehaviorByUserId(resp http.ResponseWriter, req *http.Request) {
	entityController = controllerApiCustomerBehaviorGetCustomerBehaviorByUserId.InitEntityController()
	defer entityController.Close()
	entityController.Exec(resp, req)
	return
}

package controllerWebCustomerBehavior

import (
	"github.com/kanelinweihao/lwhFrameGo/app/controller/web/customerBehavior/controllerWebCustomerBehaviorGetCustomerBehaviorByUserId"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
	"net/http"
)

var entityController typeInterface.EntityController

func GetCustomerBehaviorByUserId(resp http.ResponseWriter, req *http.Request) {
	entityController = controllerWebCustomerBehaviorGetCustomerBehaviorByUserId.InitEntityController()
	defer entityController.Close()
	entityController.Exec(resp, req)
	return
}

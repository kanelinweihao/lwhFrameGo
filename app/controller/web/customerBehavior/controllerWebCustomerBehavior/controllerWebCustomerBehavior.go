package controllerWebCustomerBehavior

import (
	"github.com/kanelinweihao/lwhFrameGo/app/controller/web/customerBehavior/controllerWebCustomerBehaviorGetCustomerBehaviorByUserId"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
	"net/http"
)

var entityController typeStruct.EntityController

func GetCustomerBehaviorByUserId(resp http.ResponseWriter, req *http.Request) {
	entityController = controllerWebCustomerBehaviorGetCustomerBehaviorByUserId.InitEntityController()
	defer entityController.Close()
	entityController.Exec(resp, req)
	return
}

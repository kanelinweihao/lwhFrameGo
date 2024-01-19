package controllerWebRate

import (
	"github.com/kanelinweihao/lwhFrameGo/app/controller/api/rate/controllerApiRateGetIRRByArrAmount"
	"github.com/kanelinweihao/lwhFrameGo/app/controller/api/rate/controllerApiRateGetRateExactlyHalf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
	"net/http"
)

var entityController typeStruct.EntityController

func GetIRRByArrAmount(resp http.ResponseWriter, req *http.Request) {
	entityController = controllerApiRateGetIRRByArrAmount.InitEntityController()
	defer entityController.Close()
	entityController.Exec(resp, req)
	return
}

func GetRateExactlyHalf(resp http.ResponseWriter, req *http.Request) {
	entityController = controllerApiRateGetRateExactlyHalf.InitEntityController()
	defer entityController.Close()
	entityController.Exec(resp, req)
	return
}

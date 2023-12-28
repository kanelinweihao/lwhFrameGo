package controllerWebRate

import (
	"github.com/kanelinweihao/lwhFrameGo/app/controller/web/rate/controllerWebRateGetIRRByArrAmount"
	"github.com/kanelinweihao/lwhFrameGo/app/controller/web/rate/controllerWebRateGetRateExactlyHalf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
	"net/http"
)

var entityController typeStruct.EntityController

func GetIRRByArrAmount(resp http.ResponseWriter, req *http.Request) {
	entityController = controllerWebRateGetIRRByArrAmount.InitEntityController()
	defer entityController.Close()
	entityController.Exec(resp, req)
	return
}

func GetRateExactlyHalf(resp http.ResponseWriter, req *http.Request) {
	entityController = controllerWebRateGetRateExactlyHalf.InitEntityController()
	defer entityController.Close()
	entityController.Exec(resp, req)
	return
}

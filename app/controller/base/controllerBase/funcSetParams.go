package controllerBase

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/middleware"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"net/http"
)

func (self *EntityControllerBase) setParams(resp http.ResponseWriter, req *http.Request) *EntityControllerBase {
	self.Resp = resp
	self.Req = req
	self.setCTX()
	self.setDefault()
	return self
}

func (self *EntityControllerBase) setCTX() *EntityControllerBase {
	req := self.Req
	jsonCTX, ok := req.Context().Value("jsonCTX").(string)
	if !ok {
		msgError := fmt.Sprintf(
			"The jsonCTX |%v| is invalid",
			jsonCTX)
		err.ErrPanic(msgError)
	}
	entityCTX := new(middleware.EntityCTX)
	entityCTX.InitByJson(jsonCTX)
	routeNameFromMiddleware := entityCTX.RouteName
	routeNameFromController := self.RouteName
	if routeNameFromMiddleware != routeNameFromController {
		msgError := fmt.Sprintf(
			"The routeName is |%s|, it should be |%s|",
			routeNameFromMiddleware,
			routeNameFromController)
		err.ErrPanic(msgError)
	}
	ipClient := entityCTX.IpClient
	self.IPClient = ipClient
	return self
}

func (self *EntityControllerBase) setDefault() *EntityControllerBase {
	entityDataController := self.Derived.GetEntityDataController()
	paramsInDefault, paramsOutDefault := entityDataController.GetParamsDefault()
	self.ParamsInDefault = paramsInDefault
	self.ParamsOutDefault = paramsOutDefault
	self.ParamsIn = self.ParamsInDefault
	self.ParamsOut = self.ParamsOutDefault
	return self
}

func (self *EntityControllerBase) GetParamsDefault() (paramsInDefault typeMap.AttrT1, paramsOutDefault typeMap.AttrT1) {
	paramsInDefault = conf.ParamsInDefaultDefault
	paramsOutDefault = conf.ParamsOutDefaultDefault
	return paramsInDefault, paramsOutDefault
}

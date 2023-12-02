package controllerBase

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/dict/dictRoute"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/funcAttr"
	"net/http"
)

func InitEntityController(resp http.ResponseWriter, req *http.Request, routeName string) (entityController *EntityController) {
	entityController = new(EntityController)
	entityController.Init(resp, req, routeName)
	return entityController
}

func (self *EntityController) Init(resp http.ResponseWriter, req *http.Request, routeName string) *EntityController {
	self.setPropertiesIn(resp, req, routeName).setPropertiesMore()
	return self
}

func (self *EntityController) setPropertiesIn(resp http.ResponseWriter, req *http.Request, routeName string) *EntityController {
	self.Resp = resp
	self.Req = req
	self.RouteName = routeName
	self.validateRoute()
	self.IsEmpty = true
	return self
}

func (self *EntityController) validateRoute() *EntityController {
	routeNameReal := self.Req.URL.Path
	attrRouteIgnoreValidate := conf.AttrRouteIgnoreValidate
	isIgnoreValidate := funcAttr.IsKeyOfAttrStr(routeNameReal, attrRouteIgnoreValidate)
	if isIgnoreValidate {
		self.IsIgnore = true
		return self
	}
	routeNameExpected := self.RouteName
	if routeNameReal != routeNameExpected {
		msgError := fmt.Sprintf(
			"The routeName is |%s|, it should be |%s|",
			routeNameReal,
			routeNameExpected)
		err.ErrPanic(msgError)
	}
	self.IsIgnore = false
	return self
}

func (self *EntityController) setPropertiesMore() *EntityController {
	if self.IsIgnore {
		return self
	}
	self.setRouteInfo().setValuesFromReq().setParamsEmpty()
	if self.IsEmpty {
		return self
	}
	self.setParamsIn()
	return self
}

func (self *EntityController) setRouteInfo() *EntityController {
	routeName := self.RouteName
	routeType, paramsInDefault, paramsOutDefault, funcService, pathTmpl := dictRoute.GetDictRouteByRouteName(routeName)
	self.RouteType = routeType
	self.ParamsInDefault = paramsInDefault
	self.ParamsOutDefault = paramsOutDefault
	self.FuncService = funcService
	self.PathTmpl = pathTmpl
	return self
}

func (self *EntityController) setValuesFromReq() *EntityController {
	valuesFromReq := self.Req.URL.Query()
	if len(valuesFromReq) == 0 {
		return self
	}
	self.ValuesFromReq = valuesFromReq
	self.IsEmpty = false
	return self
}

func (self *EntityController) setParamsEmpty() *EntityController {
	self.ParamsIn = self.ParamsInDefault
	self.ParamsOut = self.ParamsOutDefault
	return self
}

func (self *EntityController) setParamsIn() *EntityController {
	valuesFromReq := self.ValuesFromReq
	paramsInDefault := self.ParamsInDefault
	paramsIn := make(base.AttrT1)
	for field, valueDefault := range paramsInDefault {
		valueIn := valueDefault
		valueFromReq := valuesFromReq.Get(field)
		if valueFromReq != "" {
			valueIn = valueFromReq
		}
		paramsIn[field] = valueIn
	}
	self.ParamsIn = paramsIn
	return self
}

func (self *EntityController) CloseController() {
	if self.Req == nil {
		return
	}
	errClose := self.Req.Body.Close()
	err.ErrCheck(errClose)
	self.Req = nil
	self.ValuesFromReq = nil
	return
}

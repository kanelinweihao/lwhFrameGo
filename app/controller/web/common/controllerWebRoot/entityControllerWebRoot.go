package controllerWebRoot

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/controller/web/customerBehavior/controllerWebCustomerBehaviorGetCustomerBehaviorByUserId"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

var RouteType int = conf.RouteTypeWeb

type EntityController struct {
	typeStruct.EntityController
}

func (self *EntityController) GetRouteType() (routeType int) {
	routeType = RouteType
	return routeType
}

func (self *EntityController) GetParamsDefault() (paramsInDefault typeMap.AttrT1, paramsOutDefault typeMap.AttrT1) {
	paramsInDefault = controllerWebCustomerBehaviorGetCustomerBehaviorByUserId.ParamsInDefault
	paramsOutDefault = controllerWebCustomerBehaviorGetCustomerBehaviorByUserId.ParamsOutDefault
	return paramsInDefault, paramsOutDefault
}

func (self *EntityController) GetFuncService() (funcService typeStruct.FuncService) {
	funcService = controllerWebCustomerBehaviorGetCustomerBehaviorByUserId.FuncService
	return funcService
}

func (self *EntityController) GetPathTmpl() (pathTmpl string) {
	pathTmpl = controllerWebCustomerBehaviorGetCustomerBehaviorByUserId.PathTmpl
	return pathTmpl
}

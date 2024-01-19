package controllerApiRateGetIRRByArrAmount

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/service/rate/serviceRate"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

var RouteType int = conf.RouteTypeApi
var ParamsInDefault typeMap.AttrT1 = typeMap.AttrT1{
	"StrArrAmount":   "0,0,-1,-1.5,0,0,5,0,0",
	"ErrorPrecision": 4,
}
var ParamsOutDefault typeMap.AttrT1 = typeMap.AttrT1{
	"StrArrAmount":   "0,0,-1,-1.5,0,0,5,0,0",
	"ErrorPrecision": 4,
	"IRR":            "0",
}
var FuncService typeStruct.FuncService = serviceRate.GetIRRByArrAmount

type EntityController struct {
	typeStruct.EntityController
}

func (self *EntityController) GetRouteType() (routeType int) {
	routeType = RouteType
	return routeType
}

func (self *EntityController) GetParamsDefault() (paramsInDefault typeMap.AttrT1, paramsOutDefault typeMap.AttrT1) {
	paramsInDefault = ParamsInDefault
	paramsOutDefault = ParamsOutDefault
	return paramsInDefault, paramsOutDefault
}

func (self *EntityController) GetFuncService() (funcService typeStruct.FuncService) {
	funcService = FuncService
	return funcService
}

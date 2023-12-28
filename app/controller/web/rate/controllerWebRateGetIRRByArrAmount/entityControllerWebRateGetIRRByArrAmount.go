package controllerWebRateGetIRRByArrAmount

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/service/rate/serviceRate"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

var RouteType int = conf.RouteTypeWeb
var ParamsInDefault typeMap.AttrT1 = typeMap.AttrT1{
	"StrArrAmount":   "0,0,-1,-1.5,0,0,5,0,0",
	"ErrorPrecision": 4,
	"Sign":           "LessIsMore",
}
var ParamsOutDefault typeMap.AttrT1 = typeMap.AttrT1{
	"RouteNameCN":    "计算内部收益率IRR",
	"StrArrAmount":   "0,0,-1,-1.5,0,0,5,0,0",
	"ErrorPrecision": 4,
	"Sign":           "LessIsMore",
	"IRR":            "0",
}

var FuncService typeStruct.FuncService = serviceRate.GetIRRByArrAmount
var PathTmpl string = "./res/view/rate/getIRRByArrAmount.tmpl"

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

func (self *EntityController) GetPathTmpl() (pathTmpl string) {
	pathTmpl = PathTmpl
	return pathTmpl
}

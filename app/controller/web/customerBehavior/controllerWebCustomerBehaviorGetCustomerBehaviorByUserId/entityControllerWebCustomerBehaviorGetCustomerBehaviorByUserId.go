package controllerWebCustomerBehaviorGetCustomerBehaviorByUserId

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/service/customerBehavior/serviceCustomerBehavior"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

var RouteType int = conf.RouteTypeWeb
var ParamsInDefault typeMap.AttrT1 = typeMap.AttrT1{
	"UserId": 0,
	"Sign":   "LessIsMore",
}
var ParamsOutDefault typeMap.AttrT1 = typeMap.AttrT1{
	"RouteNameCN":              "获取客诉数据",
	"UserId":                   1005704,
	"Sign":                     "LessIsMore",
	"CountProductOrderPoolNFT": 0,
	"CountProductOrderNFTBuy":  0,
	"CountProductOrderNFTSell": 0,
	"SumAmountGot":             "0",
}

var FuncService typeStruct.FuncService = serviceCustomerBehavior.GetCustomerBehaviorByUserId
var PathTmpl string = "./res/view/customerBehavior/getCustomerBehaviorByUserId.tmpl"

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

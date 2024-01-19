package controllerApiUserGetMobileNoAndOrgIdByShortUserId

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/service/user/serviceUser"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

var RouteType int = conf.RouteTypeApi
var ParamsInDefault typeMap.AttrT1 = typeMap.AttrT1{
	"ShortUserId": 0,
}
var ParamsOutDefault typeMap.AttrT1 = typeMap.AttrT1{
	"ShortUserId": 1,
	"UserId":      0,
	"MobileNo":    "0",
	"OrgId":       0,
}
var FuncService typeStruct.FuncService = serviceUser.GetMobileNoAndOrgIdByShortUserId

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

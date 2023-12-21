package controllerWebUserGetUserIdByMobileNoAndOrgId

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/service/user/serviceUser"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

var RouteType int = conf.RouteTypeWeb
var ParamsInDefault typeMap.AttrT1 = typeMap.AttrT1{
	"MobileNo": 0,
	"OrgId":    0,
	"Sign":     "LessIsMore",
}
var ParamsOutDefault typeMap.AttrT1 = typeMap.AttrT1{
	"MobileNo": 13683012872,
	"OrgId":    41,
	"Sign":     "LessIsMore",
	"UserId":   0,
}
var FuncService typeStruct.FuncService = serviceUser.GetUserIdByMobileNoAndOrgId
var PathTmpl string = "./res/view/user/getUserIdByMobileNoAndOrgId.tmpl"

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

package controllerWebUserGetMobileNoAndOrgIdByShortUserId

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/service/user/serviceUser"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

var RouteType int = conf.RouteTypeWeb
var ParamsInDefault typeMap.AttrT1 = typeMap.AttrT1{
	"ShortUserId": 0,
}
var ParamsOutDefault typeMap.AttrT1 = typeMap.AttrT1{
	"RouteNameCN": "获取用户手机号和所属机构",
	"ShortUserId": 1,
	"UserId":      0,
	"MobileNo":    "0",
	"OrgId":       0,
}
var ArrEntitySectionIn []typeStruct.EntitySection = nil
var ArrEntitySectionOut []typeStruct.EntitySection = nil
var ParamsOutAppendWeb typeMap.AttrT1 = typeMap.AttrT1{
	"RouteNameCN":         "获取用户手机号和所属机构",
	"ArrEntitySectionIn":  ArrEntitySectionIn,
	"ArrEntitySectionOut": ArrEntitySectionOut,
}
var FuncService typeStruct.FuncService = serviceUser.GetMobileNoAndOrgIdByShortUserId
var PathTmpl string = "./res/view/user/getMobileNoAndOrgIdByShortUserId.tmpl"

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

func (self *EntityController) GetArrEntitySectionIn() (arrEntitySectionIn []typeStruct.EntitySection) {
	arrEntitySectionIn = ArrEntitySectionIn
	return arrEntitySectionIn
}

func (self *EntityController) GetArrEntitySectionOut() (arrEntitySectionOut []typeStruct.EntitySection) {
	arrEntitySectionOut = ArrEntitySectionOut
	return arrEntitySectionOut
}

func (self *EntityController) SetArrEntitySectionIn(arrEntitySectionIn []typeStruct.EntitySection) {
	ArrEntitySectionIn = arrEntitySectionIn
	return
}

func (self *EntityController) SetArrEntitySectionOut(arrEntitySectionOut []typeStruct.EntitySection) {
	ArrEntitySectionOut = arrEntitySectionOut
	return
}

func (self *EntityController) GetParamsOutAppendWeb() (paramsOutAppendWeb typeMap.AttrT1) {
	paramsOutAppendWeb = ParamsOutAppendWeb
	return paramsOutAppendWeb
}

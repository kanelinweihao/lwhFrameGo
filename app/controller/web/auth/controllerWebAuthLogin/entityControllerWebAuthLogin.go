package controllerWebAuthLogin

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

var RouteNameCN string = "登录"
var RouteType int = conf.RouteTypeWeb
var ParamsInDefault typeMap.AttrT1 = typeMap.AttrT1{
	"AdminUserName":     "",
	"AdminUserPassword": "",
}
var ParamsOutDefault typeMap.AttrT1 = typeMap.AttrT1{
	"AdminUserName":     "visitor",
	"AdminUserPassword": "123456",
	"AdminUserId":       2,
	"AdminUserRole":     2,
	"JwtToken":          "JwtToken",
}
var FuncService typeStruct.FuncService = nil
var PathTmpl string = "./res/view/auto/auto1.tmpl"
var ArrEntitySectionIn []typeStruct.EntitySection = []typeStruct.EntitySection{
	typeStruct.EntitySection{
		FieldName:         "AdminUserName",
		FieldNameCN:       "用户名",
		FieldRemark:       "",
		InputType:         "text",
		IsDivDisplay:      true,
		IsInputDisabled:   false,
		IsInputOnlyNumber: false,
		Value:             "",
	},
	typeStruct.EntitySection{
		FieldName:         "AdminUserPassword",
		FieldNameCN:       "密码",
		FieldRemark:       "",
		InputType:         "password",
		IsDivDisplay:      true,
		IsInputDisabled:   false,
		IsInputOnlyNumber: false,
		Value:             "",
	},
}
var ArrEntitySectionOut []typeStruct.EntitySection = []typeStruct.EntitySection{
	typeStruct.EntitySection{
		FieldName:         "AdminUserId",
		FieldNameCN:       "管理员编号",
		FieldRemark:       "",
		InputType:         "text",
		IsDivDisplay:      true,
		IsInputDisabled:   true,
		IsInputOnlyNumber: true,
		Value:             "",
	},
	typeStruct.EntitySection{
		FieldName:         "AdminUserRole",
		FieldNameCN:       "管理员角色",
		FieldRemark:       "",
		InputType:         "text",
		IsDivDisplay:      true,
		IsInputDisabled:   true,
		IsInputOnlyNumber: true,
		Value:             "",
	},
	typeStruct.EntitySection{
		FieldName:         "JwtToken",
		FieldNameCN:       "登录凭证",
		FieldRemark:       "",
		InputType:         "text",
		IsDivDisplay:      true,
		IsInputDisabled:   true,
		IsInputOnlyNumber: false,
		Value:             "",
	},
}
var TextSectionMsg string = "\n"
var IsReqToApi bool = false

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
	paramsOutAppendWeb = typeMap.AttrT1{
		"RouteNameCN":         RouteNameCN,
		"ArrEntitySectionIn":  ArrEntitySectionIn,
		"ArrEntitySectionOut": ArrEntitySectionOut,
		"TextSectionMsg":      TextSectionMsg,
		"IsReqToApi":          IsReqToApi,
	}
	return paramsOutAppendWeb
}

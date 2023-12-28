package controllerWebHome

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

var RouteType int = conf.RouteTypeWeb
var ParamsInDefault typeMap.AttrT1 = typeMap.AttrT1{}
var ParamsOutDefault typeMap.AttrT1 = typeMap.AttrT1{
	"ProjectTitle":   "未知项目名称",
	"ProjectVersion": "v1.0.0",
	"RouteNameCN":    "列表主页",
	"AttrS1Router": typeMap.AttrS1{
		"1.1 获取客诉数据":            "/web/customerBehavior/getCustomerBehaviorByUserId",
		"2.1 获取用户手机号和所属机构":      "/web/user/getMobileNoAndOrgIdByShortUserId",
		"2.2 获取UID":             "/web/user/getUserIdByMobileNoAndOrgId",
		"3.1 计算内部收益率IRR":        "/web/rate/getIRRByArrAmount",
		"3.2 计算抽中次数正好是总次数一半的概率": "/web/rate/getRateExactlyHalf",
	},
}
var FuncService typeStruct.FuncService = nil
var PathTmpl string = "./res/view/common/home.tmpl"

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

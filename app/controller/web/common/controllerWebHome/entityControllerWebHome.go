package controllerWebHome

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

var RouteType int = conf.RouteTypeWeb
var ParamsInDefault typeMap.AttrT1 = typeMap.AttrT1{}
var arrAttrS1ForSectionListRouter []typeMap.AttrS1 = []typeMap.AttrS1{
	typeMap.AttrS1{
		"RouteNameCN": "1.1 获取客诉数据",
		"RouteName":   "/web/customerBehavior/getCustomerBehaviorByUserId",
	},
	typeMap.AttrS1{
		"RouteNameCN": "2.1 获取用户手机号和所属机构",
		"RouteName":   "/web/user/getMobileNoAndOrgIdByShortUserId",
	},
	typeMap.AttrS1{
		"RouteNameCN": "2.2 获取UID",
		"RouteName":   "/web/user/getUserIdByMobileNoAndOrgId",
	},
	typeMap.AttrS1{
		"RouteNameCN": "3.1 计算内部收益率IRR",
		"RouteName":   "/web/rate/getIRRByArrAmount",
	},
	typeMap.AttrS1{
		"RouteNameCN": "3.2 计算抽中次数正好是总次数一半的概率",
		"RouteName":   "/web/rate/getRateExactlyHalf",
	},
}
var ParamsOutDefault typeMap.AttrT1 = typeMap.AttrT1{
	"RouteNameCN":                   "列表主页",
	"ArrAttrS1ForSectionListRouter": arrAttrS1ForSectionListRouter,
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

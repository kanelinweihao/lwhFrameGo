package controllerWebClientIP

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
	"RouteNameCN":    "项目版本",
}
var FuncService typeStruct.FuncService = nil
var PathTmpl string = "./res/view/common/projectInfo.tmpl"

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

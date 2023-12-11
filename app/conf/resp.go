package conf

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

const RouteTypeApi int = 1
const RouteTypeWeb int = 2

var AttrRouteNameIgnoreValidate = typeMap.AttrS1{
	"/favicon.ico": "",
	"/502.html":    "",
}

var RouteTypeDefault int = RouteTypeWeb
var ParamsInDefaultDefault typeMap.AttrT1 = typeMap.AttrT1{}
var ParamsOutDefaultDefault typeMap.AttrT1 = typeMap.AttrT1{
	"ProjectTitle":   "未知项目名称",
	"ProjectVersion": "v1.0.0",
}
var FuncServiceDefault typeStruct.FuncService = nil
var PathTmplDefault string = "./res/view/projectInfo.tmpl"

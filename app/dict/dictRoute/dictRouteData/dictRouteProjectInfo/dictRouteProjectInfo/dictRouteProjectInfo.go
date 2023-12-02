package dictRouteProjectInfo

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
)

var ParamsInDefault base.AttrT1 = base.AttrT1{}

var ParamsOutDefault base.AttrT1 = base.AttrT1{
	"ProjectTitle":   "未知项目名称",
	"ProjectVersion": "v1.0.0",
}

var FuncService conf.FuncService = nil

var PathTmpl string = "./res/view/projectInfo.tmpl"

func GetDictRoute() (paramsInDefault base.AttrT1, paramsOutDefault base.AttrT1, funcService conf.FuncService, pathTmpl string) {
	paramsInDefault = ParamsInDefault
	paramsOutDefault = ParamsOutDefault
	funcService = FuncService
	pathTmpl = PathTmpl
	return paramsInDefault, paramsOutDefault, funcService, pathTmpl
}

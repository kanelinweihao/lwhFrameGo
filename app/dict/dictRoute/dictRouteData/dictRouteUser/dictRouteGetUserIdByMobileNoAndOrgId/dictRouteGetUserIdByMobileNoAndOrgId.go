package dictRouteGetUserIdByMobileNoAndOrgId

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/service/serviceUser"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
)

var ParamsInDefault base.AttrT1 = base.AttrT1{
	"MobileNo": 0,
	"OrgId":    0,
	"Sign":     "LESSISMORE",
}

var ParamsOutDefault base.AttrT1 = base.AttrT1{
	"ProjectTitle":   "未知项目名称",
	"ProjectVersion": "v1.0.0",
	"MobileNo":       13683012872,
	"OrgId":          41,
	"Sign":           "LESSISMORE",
	"UserId":         0,
	"MsgShow":        "Ready",
}

var FuncService conf.FuncService = serviceUser.GetUserIdByMobileNoAndOrgId

var PathTmpl string = "./res/view/getUserIdByMobileNoAndOrgId.tmpl"

func GetDictRoute() (paramsInDefault base.AttrT1, paramsOutDefault base.AttrT1, funcService conf.FuncService, pathTmpl string) {
	paramsInDefault = ParamsInDefault
	paramsOutDefault = ParamsOutDefault
	funcService = FuncService
	pathTmpl = PathTmpl
	return paramsInDefault, paramsOutDefault, funcService, pathTmpl
}

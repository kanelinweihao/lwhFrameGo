package respInfoUser

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/service/serviceUser/serviceGetUserIdByMobileNoAndOrgId"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
)

var ParamsInDefaultGetUserIdByMobileNoAndOrgId base.AttrT1 = base.AttrT1{
	"MobileNo": 0,
	"OrgId":    0,
	"Sign":     "LESSISMORE",
}

var ParamsOutDefaultGetUserIdByMobileNoAndOrgId base.AttrT1 = base.AttrT1{
	"ProjectTitle":   "未知项目名称",
	"ProjectVersion": "v1.0.0",
	"MobileNo":       13683012872,
	"OrgId":          41,
	"Sign":           "LESSISMORE",
	"UserId":         0,
	"MsgShow":        "Ready",
}

var FuncServiceGetUserIdByMobileNoAndOrgId conf.FuncService = serviceGetUserIdByMobileNoAndOrgId.Exec

var PathTmplGetUserIdByMobileNoAndOrgId string = "./res/view/getUserIdByMobileNoAndOrgId.tmpl"

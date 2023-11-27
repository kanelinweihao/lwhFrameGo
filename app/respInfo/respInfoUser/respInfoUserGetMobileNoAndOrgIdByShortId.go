package respInfoUser

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/service/serviceUser/serviceGetMobileNoAndOrgIdByShortUserId"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
)

var ParamsInDefaultGetMobileNoAndOrgIdByShortUserId base.AttrT1 = base.AttrT1{
	"ShortUserId": 0,
	"Sign":        "LESSISMORE",
}

var ParamsOutDefaultGetMobileNoAndOrgIdByShortUserId base.AttrT1 = base.AttrT1{
	"ProjectTitle":   "未知项目名称",
	"ProjectVersion": "v1.0.0",
	"ShortUserId":    1,
	"Sign":           "LESSISMORE",
	"UserId":         0,
	"MobileNo":       "0",
	"OrgId":          0,
	"MsgShow":        "Ready",
}

var FuncServiceGetMobileNoAndOrgIdByShortUserId conf.FuncService = serviceGetMobileNoAndOrgIdByShortUserId.Exec

var PathTmplGetMobileNoAndOrgIdByShortUserId string = "./res/view/getMobileNoAndOrgIdByShortUserId.tmpl"

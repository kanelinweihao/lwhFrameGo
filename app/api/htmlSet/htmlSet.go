package htmlSet

import (
	"github.com/kanelinweihao/lwhFrameGo/app/api/backEnd"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/tmpl"
)

var paramsRoute = base.AttrT3{
	"/": {
		"pathTmpl": {
			"pathTmpl": "./res/view/index.tmpl",
		},
		"paramsOutDefault": base.AttrT1{
			"ProjectTitle":   "未知项目名称",
			"ProjectVersion": "v1.0.0",
			"OrgId":          "A",
			"ShortUserId":    1,
			"Sign":           "KaneIsTab",
			"MsgOut":         "Ready",
		},
		"paramsInDefault": base.AttrT1{
			"OrgId":       "AA",
			"ShortUserId": 2,
			"Sign":        "SignDefault",
		},
	},
	"/index": {
		"pathTmpl": {
			"pathTmpl": "./res/view/index.tmpl",
		},
		"paramsOutDefault": base.AttrT1{
			"ProjectTitle":   "未知项目名称",
			"ProjectVersion": "v1.0.0",
			"OrgId":          "A",
			"ShortUserId":    1,
			"Sign":           "KaneIsTab",
			"MsgOut":         "Ready",
		},
		"paramsInDefault": base.AttrT1{
			"OrgId":       "AA",
			"ShortUserId": 2,
			"Sign":        "SignDefault",
		},
	},
	"/name": {
		"pathTmpl": {
			"pathTmpl": "./res/view/name.tmpl",
		},
		"paramsOutDefault": base.AttrT1{
			"ProjectTitle":   "未知项目名称",
			"ProjectVersion": "v1.0.0",
		},
		"paramsInDefault": {},
	},
	"/version": {
		"pathTmpl": {
			"pathTmpl": "./res/view/name.tmpl",
		},
		"paramsOutDefault": base.AttrT1{
			"ProjectTitle":   "未知项目名称",
			"ProjectVersion": "v1.0.0",
		},
		"paramsInDefault": {},
	},
}

type EntityBackEnd struct{}

func (self *EntityBackEnd) ExecBackEnd(paramsIn base.AttrT1) (paramsFromBackend base.AttrT1) {
	paramsFromBackend = backEnd.ExecBackEnd(paramsIn)
	return paramsFromBackend
}

func SetHtml() {
	entityBackEnd := new(EntityBackEnd)
	tmpl.SetTmplAndServer(paramsRoute, entityBackEnd)
	return
}

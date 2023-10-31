package htmlSet

import (
	"github.com/kanelinweihao/lwhFrameGo/app/api/backEnd"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	_ "github.com/kanelinweihao/lwhFrameGo/app/utils/dd"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/tmpl"
)

var paramsRoute = base.AttrT3{
	"/": {
		"pathTmpl": {
			"pathTmpl": "./res/view/index.tmpl",
		},
		"paramsOutDefault": base.AttrT1{
			"ModTitle":   "未知项目名称",
			"ModVersion": "v1.0.0",
			"Field1":     "A",
			"Field2":     1,
			"Field3":     "NEXT",
			"MsgOut":     "Ready",
		},
		"paramsInDefault": base.AttrT1{
			"Field1": "AA",
			"Field2": 22,
			"Field3": "signHidden",
		},
	},
	"/index": {
		"pathTmpl": {
			"pathTmpl": "./res/view/index.tmpl",
		},
		"paramsOutDefault": base.AttrT1{
			"ModTitle":   "未知项目名称",
			"ModVersion": "v1.0.0",
			"Field1":     "A",
			"Field2":     1,
			"Field3":     "NEXT",
			"MsgOut":     "Ready",
		},
		"paramsInDefault": base.AttrT1{
			"Field1": "AA",
			"Field2": 22,
			"Field3": "signHidden",
		},
	},
	"/name": {
		"pathTmpl": {
			"pathTmpl": "./res/view/name.tmpl",
		},
		"paramsOutDefault": base.AttrT1{
			"ModTitle":   "未知项目名称",
			"ModVersion": "v1.0.0",
		},
		"paramsInDefault": {},
	},
	"/version": {
		"pathTmpl": {
			"pathTmpl": "./res/view/name.tmpl",
		},
		"paramsOutDefault": base.AttrT1{
			"ModTitle":   "未知项目名称",
			"ModVersion": "v1.0.0",
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

package htmlSet

import (
	"go.lwh.com/linweihao/lwhFrameGo/app/api/backEnd"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/base"
	_ "go.lwh.com/linweihao/lwhFrameGo/app/utils/dd"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/tmpl"
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
	entityBackEnd := &EntityBackEnd{}
	tmpl.SetTmplAndServer(paramsRoute, entityBackEnd)
	return
}

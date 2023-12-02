package controllerBase

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/pack"
	"html/template"
	"net/http"
	"net/url"
)

type EntityController struct {
	Resp             http.ResponseWriter
	Req              *http.Request
	IsIgnore         bool
	IsEmpty          bool
	RouteName        string
	RouteType        int
	ParamsInDefault  base.AttrT1
	ParamsOutDefault base.AttrT1
	PathTmpl         string
	FuncService      conf.FuncService
	ValuesFromReq    url.Values
	ParamsIn         base.AttrT1
	AttrT1FromSvc    base.AttrT1
	ParamsOut        base.AttrT1
	JsonRes          string
	ParamsToTmpl     base.AttrT1
	Tmpl             *template.Template
}

func (self *EntityController) Exec() *EntityController {
	if self.IsIgnore {
		return self
	}
	if !self.IsEmpty {
		self.setResFromApi().setParamsOut().setJsonRes()
	}
	routeType := self.RouteType
	switch routeType {
	case conf.RouteTypeApi:
		self.execApi()
	case conf.RouteTypeWeb:
		self.execWeb()
	default:
		msgError := fmt.Sprintf(
			"The routeType |%s| is invalid",
			routeType)
		err.ErrPanic(msgError)
	}
	return self
}

func (self *EntityController) setResFromApi() *EntityController {
	paramsIn := self.ParamsIn
	funcService := self.FuncService
	attrT1FromSvc := funcService(paramsIn)
	self.AttrT1FromSvc = attrT1FromSvc
	return self
}

func (self *EntityController) setParamsOut() *EntityController {
	paramsOutDefault := self.ParamsOutDefault
	attrT1FromSvc := self.AttrT1FromSvc
	paramsOut := make(base.AttrT1)
	for field, valueDefault := range paramsOutDefault {
		valueOut := valueDefault
		valueFromApi, ok := attrT1FromSvc[field]
		if ok {
			valueOut = valueFromApi
		}
		paramsOut[field] = valueOut
	}
	self.ParamsOut = paramsOut
	return self
}

func (self *EntityController) setJsonRes() *EntityController {
	paramsOut := self.ParamsOut
	jsonRes := getResSuccess(paramsOut)
	self.JsonRes = jsonRes
	return self
}

func (self *EntityController) execApi() *EntityController {
	self.execRes()
	return self
}

func (self *EntityController) execRes() *EntityController {
	jsonRes := self.JsonRes
	resp := self.Resp
	_, errFprintln := fmt.Fprintln(resp, jsonRes)
	err.ErrCheck(errFprintln)
	return self
}

func (self *EntityController) execWeb() *EntityController {
	self.setParamsToTmpl().setTmpl().execTmpl()
	return self
}

func (self *EntityController) setParamsToTmpl() *EntityController {
	paramsOut := self.ParamsOut
	paramsToTmpl := paramsOut
	jsonRes := self.JsonRes
	paramsToTmpl["JsonRes"] = jsonRes
	self.ParamsToTmpl = paramsToTmpl
	return self
}

func (self *EntityController) setTmpl() *EntityController {
	pathTmpl := self.PathTmpl
	fsResource, patternsTmpl := pack.GetFSAndPath(pathTmpl)
	tmpl, errTmplParse := template.ParseFS(fsResource, patternsTmpl)
	err.ErrCheck(errTmplParse)
	self.Tmpl = tmpl
	return self
}

func (self *EntityController) execTmpl() *EntityController {
	resp := self.Resp
	paramsToTmpl := self.ParamsToTmpl
	tmpl := self.Tmpl
	errTmplExec := tmpl.Execute(resp, paramsToTmpl)
	err.ErrCheck(errTmplExec)
	return self
}

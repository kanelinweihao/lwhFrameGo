package controllerBase

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/funcAttr"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/pack"
	"html/template"
)

func (self *EntityControllerBase) execWeb() *EntityControllerBase {
	self.setParamsToTmpl().setTmpl().execTmpl()
	return self
}

func (self *EntityControllerBase) setParamsToTmpl() *EntityControllerBase {
	paramsOut := self.ParamsOut
	isReqValid := self.IsReqValid
	paramsOutAppendTmpl := conf.GetParamsTmpl(isReqValid)
	paramsToTmpl := funcAttr.MergeAttr(paramsOut, paramsOutAppendTmpl)
	jsonRes := self.JsonRes
	paramsToTmpl["JsonRes"] = jsonRes
	self.ParamsToTmpl = paramsToTmpl
	return self
}

func (self *EntityControllerBase) setTmpl() *EntityControllerBase {
	pathTmpl := self.Derived.GetPathTmpl()
	self.PathTmpl = pathTmpl
	fsResource, patternsTmpl := pack.GetFSAndPath(pathTmpl)
	tmpl, errTmplParse := template.ParseFS(fsResource, patternsTmpl)
	err.ErrCheck(errTmplParse)
	self.Tmpl = tmpl
	return self
}

func (self *EntityControllerBase) GetPathTmpl() (pathTmpl string) {
	pathTmpl = conf.PathTmplDefault
	return pathTmpl
}

func (self *EntityControllerBase) execTmpl() *EntityControllerBase {
	resp := self.Resp
	paramsToTmpl := self.ParamsToTmpl
	tmpl := self.Tmpl
	errTmplExec := tmpl.Execute(resp, paramsToTmpl)
	err.ErrCheck(errTmplExec)
	return self
}

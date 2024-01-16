package controllerBase

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/file"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/funcAttr"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/pack"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
	"html/template"
)

func (self *EntityControllerBase) execWeb() *EntityControllerBase {
	self.setParamsToTmpl().setTmpl().execTmpl()
	return self
}

func (self *EntityControllerBase) setParamsToTmpl() *EntityControllerBase {
	paramsOut := self.ParamsOut
	paramsOutAppendWebBase := self.getParamsOutAppendWebBase()
	paramsOutAppendDerived := self.getParamsOutAppend()
	paramsToTmpl := funcAttr.MergeAttr(
		paramsOut,
		paramsOutAppendWebBase,
		paramsOutAppendDerived)
	self.ParamsToTmpl = paramsToTmpl
	// dd.DD(self.ParamsToTmpl)
	return self
}

func (self *EntityControllerBase) getParamsOutAppendWebBase() (paramsAppend typeMap.AttrT1) {
	projectTitle := conf.GetProjectNameCN()
	projectVersion := conf.GetProjectVersion()
	ipClient := self.IPClient
	isReqValid := self.IsReqValid
	msgShow := "Ready"
	if isReqValid {
		msgShow = "Success"
	}
	jsonRes := self.JsonRes
	paramsAppend = typeMap.AttrT1{
		"ProjectTitle":   projectTitle,
		"ProjectVersion": projectVersion,
		"ClientIP":       ipClient,
		"MsgShow":        msgShow,
		"JsonRes":        jsonRes,
	}
	return paramsAppend
}

func (self *EntityControllerBase) getParamsOutAppend() (paramsOutAppendWeb typeMap.AttrT1) {
	arrEntitySectionIn := self.getArrEntitySectionIn()
	self.Derived.SetArrEntitySectionIn(arrEntitySectionIn)
	arrEntitySectionOut := self.getArrEntitySectionOut()
	self.Derived.SetArrEntitySectionOut(arrEntitySectionOut)
	paramsOutAppendWeb = self.Derived.GetParamsOutAppendWeb()
	return paramsOutAppendWeb
}

func (self *EntityControllerBase) getArrEntitySectionIn() (arrEntitySectionIn []typeStruct.EntitySection) {
	paramsOut := self.GetParamsOut()
	arrEntitySectionIn = self.Derived.GetArrEntitySectionIn()
	arrEntitySectionIn = getArrEntitySectionWithValue(arrEntitySectionIn, paramsOut)
	return arrEntitySectionIn
}

func (self *EntityControllerBase) getArrEntitySectionOut() (arrEntitySectionOut []typeStruct.EntitySection) {
	paramsOut := self.GetParamsOut()
	arrEntitySectionOut = self.Derived.GetArrEntitySectionOut()
	arrEntitySectionOut = getArrEntitySectionWithValue(arrEntitySectionOut, paramsOut)
	return arrEntitySectionOut
}

func getArrEntitySectionWithValue(arrEntitySectionWithoutValue []typeStruct.EntitySection, paramsValue typeMap.AttrT1) (arrEntitySection []typeStruct.EntitySection) {
	arrEntitySection = make([]typeStruct.EntitySection, len(arrEntitySectionWithoutValue))
	for k, entitySection := range arrEntitySectionWithoutValue {
		fieldName := entitySection.FieldName
		value := conv.ToStr(paramsValue[fieldName])
		entitySection.Value = value
		arrEntitySection[k] = entitySection
	}
	return arrEntitySection
}

func (self *EntityControllerBase) GetArrEntitySectionIn() (arrEntitySectionIn []typeStruct.EntitySection) {
	arrEntitySectionIn = conf.ArrEntitySectionIn
	return arrEntitySectionIn
}

func (self *EntityControllerBase) GetArrEntitySectionOut() (arrEntitySectionOut []typeStruct.EntitySection) {
	arrEntitySectionOut = conf.ArrEntitySectionOut
	return arrEntitySectionOut
}

func (self *EntityControllerBase) SetArrEntitySectionIn(arrEntitySectionIn []typeStruct.EntitySection) {
	return
}

func (self *EntityControllerBase) SetArrEntitySectionOut(arrEntitySectionOut []typeStruct.EntitySection) {
	return
}

func (self *EntityControllerBase) GetParamsOutAppendWeb() (paramsOutAppendWeb typeMap.AttrT1) {
	paramsOutAppendWeb = conf.ParamsOutAppendWeb
	return paramsOutAppendWeb
}

func (self *EntityControllerBase) setTmpl() *EntityControllerBase {
	pathTmpl := self.Derived.GetPathTmpl()
	self.PathTmpl = pathTmpl
	fsResource, pathEmbedTmpl := pack.GetFSAndPath(pathTmpl)
	fileName := file.GetFilename(pathEmbedTmpl)
	tmpl := template.New(fileName)
	tmplFuncMap := conf.GetTmplFuncMap()
	tmpl.Funcs(tmplFuncMap)
	pathViewHeader := conf.GetPathViewHeader()
	pathViewFooter := conf.GetPathViewFooter()
	pathViewSectionProject := conf.GetPathViewSectionProject()
	pathViewJSSM3 := conf.GetPathViewJSSM3()
	pathViewJSSubmitReq := conf.GetPathViewJSSubmitReq()
	tmpl, errTmplParse := tmpl.ParseFS(
		fsResource,
		pathEmbedTmpl,
		pathViewHeader,
		pathViewFooter,
		pathViewSectionProject,
		pathViewJSSM3,
		pathViewJSSubmitReq)
	tmpl = template.Must(tmpl, errTmplParse)
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

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
	return self
}

func (self *EntityControllerBase) getParamsOutAppendWebBase() (paramsAppend typeMap.AttrT1) {
	projectTitle := conf.GetProjectNameCN()
	projectVersion := conf.GetProjectVersion()
	ipClient := self.IPClient
	isRespDefault := self.IsRespDefault
	msgShow := "Ready"
	if !isRespDefault {
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
	self.setArrEntitySectionInToDerived()
	self.setArrEntitySectionOutToDerived()
	entityDataController := self.Derived.GetEntityDataController()
	paramsOutAppendWeb = entityDataController.GetParamsOutAppendWeb()
	return paramsOutAppendWeb
}

func (self *EntityControllerBase) setArrEntitySectionInToDerived() *EntityControllerBase {
	paramsOut := self.ParamsOut
	entityDataController := self.Derived.GetEntityDataController()
	arrEntitySectionIn := entityDataController.GetArrEntitySectionIn()
	arrEntitySectionIn = getArrEntitySectionWithValue(arrEntitySectionIn, paramsOut)
	entityDataController.SetArrEntitySectionIn(arrEntitySectionIn)
	return self
}

func (self *EntityControllerBase) setArrEntitySectionOutToDerived() *EntityControllerBase {
	paramsOut := self.ParamsOut
	entityDataController := self.Derived.GetEntityDataController()
	arrEntitySectionOut := entityDataController.GetArrEntitySectionOut()
	arrEntitySectionOut = getArrEntitySectionWithValue(arrEntitySectionOut, paramsOut)
	entityDataController.SetArrEntitySectionOut(arrEntitySectionOut)
	return self
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
	entityDataController := self.Derived.GetEntityDataController()
	pathTmpl := entityDataController.GetPathTmpl()
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
	pathViewJSHome := conf.GetPathViewJSHome()
	tmpl, errTmplParse := tmpl.ParseFS(
		fsResource,
		pathEmbedTmpl,
		pathViewHeader,
		pathViewFooter,
		pathViewSectionProject,
		pathViewJSSM3,
		pathViewJSSubmitReq,
		pathViewJSHome)
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

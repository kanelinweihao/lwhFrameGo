package controllerBase

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/logs"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/res"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

func (self *EntityControllerBase) setResDefault() *EntityControllerBase {
	self.setParamsDefault().setJsonRes()
	return self
}

func (self *EntityControllerBase) setRes() *EntityControllerBase {
	self.setParamsDefault().setParamsIn().setParamsFromSvc().setParamsOut().setJsonRes()
	return self
}

func (self *EntityControllerBase) setParamsDefault() *EntityControllerBase {
	paramsInDefault, paramsOutDefault := self.Derived.GetParamsDefault()
	self.ParamsInDefault = paramsInDefault
	self.ParamsOutDefault = paramsOutDefault
	self.ParamsIn = self.ParamsInDefault
	self.ParamsOut = self.ParamsOutDefault
	return self
}

func (self *EntityControllerBase) GetParamsDefault() (paramsInDefault typeMap.AttrT1, paramsOutDefault typeMap.AttrT1) {
	paramsInDefault = conf.ParamsInDefaultDefault
	paramsOutDefault = conf.ParamsOutDefaultDefault
	return paramsInDefault, paramsOutDefault
}

func (self *EntityControllerBase) setParamsIn() *EntityControllerBase {
	valuesFromReq := self.ValuesFromReq
	paramsInDefault := self.ParamsInDefault
	paramsIn := make(typeMap.AttrT1)
	for field, valueDefault := range paramsInDefault {
		valueIn := valueDefault
		valueFromReq := valuesFromReq.Get(field)
		if valueFromReq != "" {
			valueIn = valueFromReq
		}
		paramsIn[field] = valueIn
	}
	self.ParamsIn = paramsIn
	return self
}

func (self *EntityControllerBase) setParamsFromSvc() *EntityControllerBase {
	paramsIn := self.ParamsIn
	funcService := self.Derived.GetFuncService()
	if funcService == nil {
		self.ParamsFromSvc = make(typeMap.AttrT1)
		return self
	}
	self.FuncService = funcService
	paramsFromSvc := funcService(paramsIn)
	self.ParamsFromSvc = paramsFromSvc
	return self
}

func (self *EntityControllerBase) GetFuncService() (funcService typeStruct.FuncService) {
	funcService = conf.FuncServiceDefault
	return funcService
}

func (self *EntityControllerBase) setParamsOut() *EntityControllerBase {
	paramsOutDefault := self.ParamsOutDefault
	paramsFromSvc := self.ParamsFromSvc
	paramsOut := make(typeMap.AttrT1)
	for field, valueDefault := range paramsOutDefault {
		valueOut := valueDefault
		valueFromSvc, ok := paramsFromSvc[field]
		if ok {
			valueOut = valueFromSvc
		}
		paramsOut[field] = valueOut
	}
	self.ParamsOut = paramsOut
	return self
}

func (self *EntityControllerBase) setJsonRes() *EntityControllerBase {
	paramsOut := self.ParamsOut
	jsonRes := res.GetResSuccess(paramsOut)
	self.JsonRes = jsonRes
	logs.SetLog(jsonRes)
	return self
}

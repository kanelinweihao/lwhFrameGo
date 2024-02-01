package controllerBase

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/logs"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/res"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeFunc"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
)

func (self *EntityControllerBase) execSvc() *EntityControllerBase {
	isRespDefault := self.isRespDefault()
	if !isRespDefault {
		self.setParamsIn().setParamsFromSvc().setParamsOut()
	}
	self.setJsonRes()
	return self
}

func (self *EntityControllerBase) isRespDefault() (isRespDefault bool) {
	req := self.Req
	countValues := len(req.URL.Query())
	isRespDefault = countValues == 0
	self.IsRespDefault = isRespDefault
	return isRespDefault
}

func (self *EntityControllerBase) setParamsIn() *EntityControllerBase {
	req := self.Req
	valuesFromReq := req.URL.Query()
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
	entityDataController := self.Derived.GetEntityDataController()
	funcService := entityDataController.GetFuncService()
	if funcService == nil {
		self.ParamsFromSvc = make(typeMap.AttrT1)
		return self
	}
	self.FuncService = funcService
	paramsFromSvc := funcService(paramsIn)
	self.ParamsFromSvc = paramsFromSvc
	return self
}

func (self *EntityControllerBase) GetFuncService() (funcService typeFunc.FuncService) {
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
	logs.SetLog(jsonRes, "res")
	return self
}

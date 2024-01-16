package controllerWebRateGetIRRByArrAmount

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/service/rate/serviceRate"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

var RouteType int = conf.RouteTypeWeb
var ParamsInDefault typeMap.AttrT1 = typeMap.AttrT1{
	"StrArrAmount":   "0,0,-1,-1.5,0,0,5,0,0",
	"ErrorPrecision": 4,
}
var ParamsOutDefault typeMap.AttrT1 = typeMap.AttrT1{
	"StrArrAmount":   "0,0,-1,-1.5,0,0,5,0,0",
	"ErrorPrecision": 4,
	"IRR":            "0",
}
var ArrEntitySectionIn []typeStruct.EntitySection = []typeStruct.EntitySection{
	typeStruct.EntitySection{
		FieldName:     "StrArrAmount",
		FieldNameCN:   "每期变化值",
		FieldRemark:   "用英文逗号隔开,负数为支出,正数为收入",
		Value:         "",
		InputType:     "text",
		DivDisplay:    "TRUE",
		InputDisabled: "FALSE",
		InputOnkeyup:  "onlyNumber",
	},
	typeStruct.EntitySection{
		FieldName:     "ErrorPrecision",
		FieldNameCN:   "误差精度",
		FieldRemark:   "IRR小数表示时,小数点后保留位数",
		Value:         "",
		InputType:     "text",
		DivDisplay:    "TRUE",
		InputDisabled: "TRUE",
		InputOnkeyup:  "onlyNumber",
	},
}
var ArrEntitySectionOut []typeStruct.EntitySection = []typeStruct.EntitySection{
	typeStruct.EntitySection{
		FieldName:     "IRR",
		FieldNameCN:   "内部收益率",
		FieldRemark:   "",
		Value:         "",
		InputType:     "text",
		DivDisplay:    "TRUE",
		InputDisabled: "FALSE",
		InputOnkeyup:  "onlyNumber",
	},
}
var ParamsOutAppendWeb typeMap.AttrT1 = typeMap.AttrT1{
	"RouteNameCN":         "计算内部收益率IRR",
	"ArrEntitySectionIn":  ArrEntitySectionIn,
	"ArrEntitySectionOut": ArrEntitySectionOut,
}
var FuncService typeStruct.FuncService = serviceRate.GetIRRByArrAmount
var PathTmpl string = "./res/view/rate/getIRRByArrAmount.tmpl"

type EntityController struct {
	typeStruct.EntityController
}

func (self *EntityController) GetRouteType() (routeType int) {
	routeType = RouteType
	return routeType
}

func (self *EntityController) GetParamsDefault() (paramsInDefault typeMap.AttrT1, paramsOutDefault typeMap.AttrT1) {
	paramsInDefault = ParamsInDefault
	paramsOutDefault = ParamsOutDefault
	return paramsInDefault, paramsOutDefault
}

func (self *EntityController) GetFuncService() (funcService typeStruct.FuncService) {
	funcService = FuncService
	return funcService
}

func (self *EntityController) GetPathTmpl() (pathTmpl string) {
	pathTmpl = PathTmpl
	return pathTmpl
}

func (self *EntityController) GetArrEntitySectionIn() (arrEntitySectionIn []typeStruct.EntitySection) {
	arrEntitySectionIn = ArrEntitySectionIn
	return arrEntitySectionIn
}

func (self *EntityController) GetArrEntitySectionOut() (arrEntitySectionOut []typeStruct.EntitySection) {
	arrEntitySectionOut = ArrEntitySectionOut
	return arrEntitySectionOut
}

func (self *EntityController) SetArrEntitySectionIn(arrEntitySectionIn []typeStruct.EntitySection) {
	ArrEntitySectionIn = arrEntitySectionIn
	return
}

func (self *EntityController) SetArrEntitySectionOut(arrEntitySectionOut []typeStruct.EntitySection) {
	ArrEntitySectionOut = arrEntitySectionOut
	return
}

func (self *EntityController) GetParamsOutAppendWeb() (paramsOutAppendWeb typeMap.AttrT1) {
	paramsOutAppendWeb = ParamsOutAppendWeb
	return paramsOutAppendWeb
}

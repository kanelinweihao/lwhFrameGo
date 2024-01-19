package controllerWebRateGetRateExactlyHalf

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/service/rate/serviceRate"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

var RouteNameCN string = "计算抽中次数正好是总次数一半的概率"
var RouteType int = conf.RouteTypeWeb
var ParamsInDefault typeMap.AttrT1 = typeMap.AttrT1{
	"X": 10,
}
var ParamsOutDefault typeMap.AttrT1 = typeMap.AttrT1{
	"X":    10,
	"Rate": "0",
}
var FuncService typeStruct.FuncService = serviceRate.GetRateExactlyHalf
var PathTmpl string = "./res/view/auto/auto1.tmpl"
var ArrEntitySectionIn []typeStruct.EntitySection = []typeStruct.EntitySection{
	typeStruct.EntitySection{
		FieldName:         "X",
		FieldNameCN:       "总抽奖次数",
		FieldRemark:       "",
		InputType:         "text",
		IsDivDisplay:      true,
		IsInputDisabled:   false,
		IsInputOnlyNumber: true,
		Value:             "",
	},
}
var ArrEntitySectionOut []typeStruct.EntitySection = []typeStruct.EntitySection{
	typeStruct.EntitySection{
		FieldName:         "Rate",
		FieldNameCN:       "目标概率",
		FieldRemark:       "",
		InputType:         "text",
		IsDivDisplay:      true,
		IsInputDisabled:   true,
		IsInputOnlyNumber: true,
		Value:             "",
	},
}
var TextSectionMsg string = "假设每次抽奖的中奖概率为[50%](即一半概率抽中一半概率抽不中),\n若总共抽奖[X]次,求抽中次数正好为总次数的一半的概率[Rate]\n"
var IsReqToApi bool = false

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
	paramsOutAppendWeb = typeMap.AttrT1{
		"RouteNameCN":         RouteNameCN,
		"ArrEntitySectionIn":  ArrEntitySectionIn,
		"ArrEntitySectionOut": ArrEntitySectionOut,
		"TextSectionMsg":      TextSectionMsg,
		"IsReqToApi":          IsReqToApi,
	}
	return paramsOutAppendWeb
}

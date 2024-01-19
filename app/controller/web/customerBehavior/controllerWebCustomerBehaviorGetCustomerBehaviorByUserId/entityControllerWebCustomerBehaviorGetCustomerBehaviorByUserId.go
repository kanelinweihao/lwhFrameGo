package controllerWebCustomerBehaviorGetCustomerBehaviorByUserId

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/service/customerBehavior/serviceCustomerBehavior"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

var RouteNameCN string = "获取客诉数据"
var RouteType int = conf.RouteTypeWeb
var ParamsInDefault typeMap.AttrT1 = typeMap.AttrT1{
	"UserId": 0,
}
var ParamsOutDefault typeMap.AttrT1 = typeMap.AttrT1{
	"UserId":                   1005704,
	"CountProductOrderPoolNFT": 0,
	"CountProductOrderNFTBuy":  0,
	"CountProductOrderNFTSell": 0,
	"SumAmountGot":             "0",
}
var FuncService typeStruct.FuncService = serviceCustomerBehavior.GetCustomerBehaviorByUserId
var PathTmpl string = "./res/view/auto/auto1.tmpl"
var ArrEntitySectionIn []typeStruct.EntitySection = []typeStruct.EntitySection{
	typeStruct.EntitySection{
		FieldName:         "UserId",
		FieldNameCN:       "用户编号",
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
		FieldName:         "CountProductOrderPoolNFT",
		FieldNameCN:       "首发买入条数",
		FieldRemark:       "",
		InputType:         "text",
		IsDivDisplay:      true,
		IsInputDisabled:   true,
		IsInputOnlyNumber: true,
		Value:             "",
	},
	typeStruct.EntitySection{
		FieldName:         "CountProductOrderNFTBuy",
		FieldNameCN:       "寄售买入条数",
		FieldRemark:       "",
		InputType:         "text",
		IsDivDisplay:      true,
		IsInputDisabled:   true,
		IsInputOnlyNumber: true,
		Value:             "",
	},
	typeStruct.EntitySection{
		FieldName:         "CountProductOrderNFTSell",
		FieldNameCN:       "寄售卖出条数",
		FieldRemark:       "",
		InputType:         "text",
		IsDivDisplay:      true,
		IsInputDisabled:   true,
		IsInputOnlyNumber: true,
		Value:             "",
	},
	typeStruct.EntitySection{
		FieldName:         "SumAmountGot",
		FieldNameCN:       "用户水晶已领取总数",
		FieldRemark:       "",
		InputType:         "text",
		IsDivDisplay:      true,
		IsInputDisabled:   true,
		IsInputOnlyNumber: true,
		Value:             "",
	},
}
var TextSectionMsg string = "\n"
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

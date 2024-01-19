package controllerWebRoot

import (
	"github.com/kanelinweihao/lwhFrameGo/app/controller/web/customerBehavior/controllerWebCustomerBehaviorGetCustomerBehaviorByUserId"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

type EntityController struct {
	typeStruct.EntityController
}

func (self *EntityController) GetRouteType() (routeType int) {
	routeType = controllerWebCustomerBehaviorGetCustomerBehaviorByUserId.RouteType
	return routeType
}

func (self *EntityController) GetParamsDefault() (paramsInDefault typeMap.AttrT1, paramsOutDefault typeMap.AttrT1) {
	paramsInDefault = controllerWebCustomerBehaviorGetCustomerBehaviorByUserId.ParamsInDefault
	paramsOutDefault = controllerWebCustomerBehaviorGetCustomerBehaviorByUserId.ParamsOutDefault
	return paramsInDefault, paramsOutDefault
}

func (self *EntityController) GetFuncService() (funcService typeStruct.FuncService) {
	funcService = controllerWebCustomerBehaviorGetCustomerBehaviorByUserId.FuncService
	return funcService
}

func (self *EntityController) GetPathTmpl() (pathTmpl string) {
	pathTmpl = controllerWebCustomerBehaviorGetCustomerBehaviorByUserId.PathTmpl
	return pathTmpl
}

func (self *EntityController) GetArrEntitySectionIn() (arrEntitySectionIn []typeStruct.EntitySection) {
	arrEntitySectionIn = controllerWebCustomerBehaviorGetCustomerBehaviorByUserId.ArrEntitySectionIn
	return arrEntitySectionIn
}

func (self *EntityController) GetArrEntitySectionOut() (arrEntitySectionOut []typeStruct.EntitySection) {
	arrEntitySectionOut = controllerWebCustomerBehaviorGetCustomerBehaviorByUserId.ArrEntitySectionOut
	return arrEntitySectionOut
}

func (self *EntityController) SetArrEntitySectionIn(arrEntitySectionIn []typeStruct.EntitySection) {
	controllerWebCustomerBehaviorGetCustomerBehaviorByUserId.ArrEntitySectionIn = arrEntitySectionIn
	return
}

func (self *EntityController) SetArrEntitySectionOut(arrEntitySectionOut []typeStruct.EntitySection) {
	controllerWebCustomerBehaviorGetCustomerBehaviorByUserId.ArrEntitySectionOut = arrEntitySectionOut
	return
}

func (self *EntityController) GetParamsOutAppendWeb() (paramsOutAppendWeb typeMap.AttrT1) {
	paramsOutAppendWeb = typeMap.AttrT1{
		"RouteNameCN":         controllerWebCustomerBehaviorGetCustomerBehaviorByUserId.RouteNameCN,
		"ArrEntitySectionIn":  controllerWebCustomerBehaviorGetCustomerBehaviorByUserId.ArrEntitySectionIn,
		"ArrEntitySectionOut": controllerWebCustomerBehaviorGetCustomerBehaviorByUserId.ArrEntitySectionOut,
		"TextSectionMsg":      controllerWebCustomerBehaviorGetCustomerBehaviorByUserId.TextSectionMsg,
		"IsReqToApi":          controllerWebCustomerBehaviorGetCustomerBehaviorByUserId.IsReqToApi,
	}
	return paramsOutAppendWeb
}

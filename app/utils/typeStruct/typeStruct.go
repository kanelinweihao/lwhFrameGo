package typeStruct

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/consts"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/funcStruct"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeFunc"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"strings"
)

type EntityDataController struct {
	Basic EntityDataControllerBasic
	Web   EntityDataControllerWeb
}

type EntityDataControllerBasic struct {
	RouteName        string
	RouteType        int
	FuncService      typeFunc.FuncService
	ParamsInDefault  typeMap.AttrT1
	ParamsOutDefault typeMap.AttrT1
}

type EntityDataControllerWeb struct {
	PathTmpl            string
	IsReqToApi          bool
	RouteNameCN         string
	TextSectionMsg      string
	ArrEntitySectionIn  []EntitySection
	ArrEntitySectionOut []EntitySection
}

type EntitySection struct {
	FieldName         string `validate:"required,min=2"`
	FieldNameCN       string `validate:"required,min=2"`
	FieldRemark       string `validate:"required,min=2"`
	InputType         string `validate:"required,oneof='text' 'password'"`
	IsDivDisplay      bool   `validate:"required"`
	IsInputDisabled   bool   `validate:"required"`
	IsInputOnlyNumber bool   `validate:"required"`
	Value             string `validate:"omitempty"`
}

func (self *EntityDataController) ToEntityAPI() (entityAPI *EntityDataController) {
	entityAPI = new(EntityDataController)
	funcStruct.StructCopy(entityAPI, self)
	entityAPI.Basic.RouteType = consts.RouteTypeApi
	entityAPI.Basic.RouteName = strings.Replace(
		entityAPI.Basic.RouteName,
		"web",
		"api",
		1)
	return entityAPI
}

func (self *EntityDataController) GetRouteType() (routeType int) {
	routeType = self.Basic.RouteType
	return routeType
}

func (self *EntityDataController) GetParamsDefault() (paramsInDefault typeMap.AttrT1, paramsOutDefault typeMap.AttrT1) {
	paramsInDefault = self.Basic.ParamsInDefault
	paramsOutDefault = self.Basic.ParamsOutDefault
	return paramsInDefault, paramsOutDefault
}

func (self *EntityDataController) GetFuncService() (funcService typeFunc.FuncService) {
	funcService = self.Basic.FuncService
	return funcService
}

func (self *EntityDataController) GetPathTmpl() (pathTmpl string) {
	pathTmpl = self.Web.PathTmpl
	return pathTmpl
}

func (self *EntityDataController) GetArrEntitySectionIn() (arrEntitySectionIn []EntitySection) {
	arrEntitySectionIn = self.Web.ArrEntitySectionIn
	return arrEntitySectionIn
}

func (self *EntityDataController) GetArrEntitySectionOut() (arrEntitySectionOut []EntitySection) {
	arrEntitySectionOut = self.Web.ArrEntitySectionOut
	return arrEntitySectionOut
}

func (self *EntityDataController) SetArrEntitySectionIn(arrEntitySectionIn []EntitySection) {
	self.Web.ArrEntitySectionIn = arrEntitySectionIn
	return
}

func (self *EntityDataController) SetArrEntitySectionOut(arrEntitySectionOut []EntitySection) {
	self.Web.ArrEntitySectionOut = arrEntitySectionOut
	return
}

func (self *EntityDataController) GetParamsOutAppendWeb() (paramsOutAppendWeb typeMap.AttrT1) {
	paramsOutAppendWeb = typeMap.AttrT1{
		"IsReqToApi":          self.Web.IsReqToApi,
		"RouteNameCN":         self.Web.RouteNameCN,
		"TextSectionMsg":      self.Web.TextSectionMsg,
		"ArrEntitySectionIn":  self.Web.ArrEntitySectionIn,
		"ArrEntitySectionOut": self.Web.ArrEntitySectionOut,
	}
	return paramsOutAppendWeb
}

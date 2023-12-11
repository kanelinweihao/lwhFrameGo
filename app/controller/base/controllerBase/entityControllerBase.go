package controllerBase

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/funcAttr"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/times"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
	"html/template"
	"net/http"
	"net/url"
)

type EntityControllerBase struct {
	Derived          typeStruct.EntityController
	RouteName        string
	Resp             http.ResponseWriter
	Req              *http.Request
	IsReqValid       bool
	ParamsInDefault  typeMap.AttrT1
	ParamsOutDefault typeMap.AttrT1
	ValuesFromReq    url.Values
	ParamsFromSvc    typeMap.AttrT1
	FuncService      typeStruct.FuncService
	ParamsIn         typeMap.AttrT1
	ParamsOut        typeMap.AttrT1
	JsonRes          string
	RouteType        int
	ParamsToTmpl     typeMap.AttrT1
	PathTmpl         string
	Tmpl             *template.Template
}

func (self *EntityControllerBase) Exec(resp http.ResponseWriter, req *http.Request) {
	self.Resp = resp
	self.Req = req
	isRouteNameIgnore := self.isRouteNameIgnore()
	if isRouteNameIgnore {
		self.IsReqValid = false
		return
	}
	isReqValuesEmpty := self.isReqValuesEmpty()
	if isReqValuesEmpty {
		self.IsReqValid = false
		self.setResDefault()
	} else {
		self.IsReqValid = true
		self.setRes()
	}
	self.execFront()
	times.ShowTimeAndMsg("OK")
	return
}

func (self *EntityControllerBase) isRouteNameIgnore() (isRouteNameIgnore bool) {
	routeNameReal := self.Req.URL.Path
	attrRouteNameIgnoreValidate := conf.AttrRouteNameIgnoreValidate
	isRouteNameIgnore = funcAttr.IsKeyOfAttrStr(routeNameReal, attrRouteNameIgnoreValidate)
	if isRouteNameIgnore {
		return true
	}
	routeNameExpected := self.RouteName
	if routeNameReal != routeNameExpected {
		msgError := fmt.Sprintf(
			"The routeName is |%s|, it should be |%s|",
			routeNameReal,
			routeNameExpected)
		err.ErrPanic(msgError)
	}
	return false
}

func (self *EntityControllerBase) isReqValuesEmpty() (isReqValuesEmpty bool) {
	valuesFromReq := self.Req.URL.Query()
	if len(valuesFromReq) == 0 {
		self.ValuesFromReq = nil
		return true
	}
	self.ValuesFromReq = valuesFromReq
	return false
}

func (self *EntityControllerBase) execFront() {
	routeType := self.Derived.GetRouteType()
	self.RouteType = routeType
	switch routeType {
	case conf.RouteTypeApi:
		self.execApi()
	case conf.RouteTypeWeb:
		self.execWeb()
	default:
		msgError := fmt.Sprintf(
			"The routeType |%s| is invalid",
			routeType)
		err.ErrPanic(msgError)
	}
	return
}

func (self *EntityControllerBase) GetRouteType() (routeType int) {
	routeType = conf.RouteTypeDefault
	return routeType
}

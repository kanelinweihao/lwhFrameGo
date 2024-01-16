package controllerBase

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/ip"
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
	IPClient         string
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
	self.checkRouteNameValid()
	self.setIPOfClient()
	isReqValid := self.isReqValid()
	if isReqValid {
		self.setRes()
	} else {
		self.setResDefault()
	}
	self.execFront()
	return
}

func (self *EntityControllerBase) checkRouteNameValid() {
	routeNameReal := self.Req.URL.Path
	routeNameExpected := self.RouteName
	if routeNameReal != routeNameExpected {
		msgError := fmt.Sprintf(
			"The routeName is |%s|, it should be |%s|",
			routeNameReal,
			routeNameExpected)
		err.ErrPanic(msgError)
	}
	return
}

func (self *EntityControllerBase) setIPOfClient() {
	req := self.Req
	ipClient := ip.GetIPOfClient(req)
	self.IPClient = ipClient
	return
}

func (self *EntityControllerBase) isReqValid() (isReqValid bool) {
	valuesFromReq := self.Req.URL.Query()
	if len(valuesFromReq) == 0 {
		self.ValuesFromReq = nil
		self.IsReqValid = false
		return self.IsReqValid
	}
	self.ValuesFromReq = valuesFromReq
	self.IsReqValid = true
	return self.IsReqValid
}

func (self *EntityControllerBase) execFront() {
	routeType := self.Derived.GetRouteType()
	self.RouteType = routeType
	switch routeType {
	case conf.RouteTypeFile:
		self.execFile()
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

func (self *EntityControllerBase) GetParamsOut() (paramsOut typeMap.AttrT1) {
	paramsOut = self.ParamsOut
	return paramsOut
}

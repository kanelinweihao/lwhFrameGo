package controllerBase

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
	"html/template"
	"net/http"
)

type EntityControllerBase struct {
	Derived          typeStruct.EntityController
	RouteName        string
	Resp             http.ResponseWriter
	Req              *http.Request
	IPClient         string
	IsRespDefault    bool
	ParamsInDefault  typeMap.AttrT1
	ParamsOutDefault typeMap.AttrT1
	FuncService      typeStruct.FuncService
	ParamsFromSvc    typeMap.AttrT1
	ParamsIn         typeMap.AttrT1
	ParamsOut        typeMap.AttrT1
	JsonRes          string
	RouteType        int
	ParamsToTmpl     typeMap.AttrT1
	PathTmpl         string
	Tmpl             *template.Template
}

func (self *EntityControllerBase) Exec(resp http.ResponseWriter, req *http.Request) {
	self.setParams(resp, req).execSvc().execFront()
	return
}

func (self *EntityControllerBase) GetIpClient() (ipClient string) {
	ipClient = self.IPClient
	return ipClient
}

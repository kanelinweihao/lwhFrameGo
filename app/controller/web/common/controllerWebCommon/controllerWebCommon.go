package controllerWebCommon

import (
	"github.com/kanelinweihao/lwhFrameGo/app/controller/web/common/controllerWebClientIP"
	"github.com/kanelinweihao/lwhFrameGo/app/controller/web/common/controllerWebHome"
	"github.com/kanelinweihao/lwhFrameGo/app/controller/web/common/controllerWebProjectTitle"
	"github.com/kanelinweihao/lwhFrameGo/app/controller/web/common/controllerWebProjectVersion"
	"github.com/kanelinweihao/lwhFrameGo/app/controller/web/common/controllerWebRoot"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
	"net/http"
)

var entityController typeStruct.EntityController

func Root(resp http.ResponseWriter, req *http.Request) {
	entityController = controllerWebRoot.InitEntityController()
	defer entityController.Close()
	entityController.Exec(resp, req)
	return
}

func Home(resp http.ResponseWriter, req *http.Request) {
	entityController = controllerWebHome.InitEntityController()
	defer entityController.Close()
	entityController.Exec(resp, req)
	return
}

func ProjectTitle(resp http.ResponseWriter, req *http.Request) {
	entityController = controllerWebProjectTitle.InitEntityController()
	defer entityController.Close()
	entityController.Exec(resp, req)
	return
}

func ProjectVersion(resp http.ResponseWriter, req *http.Request) {
	entityController = controllerWebProjectVersion.InitEntityController()
	defer entityController.Close()
	entityController.Exec(resp, req)
	return
}

func ClientIP(resp http.ResponseWriter, req *http.Request) {
	entityController = controllerWebClientIP.InitEntityController()
	defer entityController.Close()
	entityController.Exec(resp, req)
	return
}

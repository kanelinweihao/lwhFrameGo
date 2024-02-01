package controllerApiCommon

import (
	"github.com/kanelinweihao/lwhFrameGo/app/controller/api/common/controllerApiClientIP"
	"github.com/kanelinweihao/lwhFrameGo/app/controller/api/common/controllerApiProjectTitle"
	"github.com/kanelinweihao/lwhFrameGo/app/controller/api/common/controllerApiProjectVersion"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
	"net/http"
)

var entityController typeInterface.EntityController

func ProjectTitle(resp http.ResponseWriter, req *http.Request) {
	entityController = controllerApiProjectTitle.InitEntityController()
	defer entityController.Close()
	entityController.Exec(resp, req)
	return
}

func ProjectVersion(resp http.ResponseWriter, req *http.Request) {
	entityController = controllerApiProjectVersion.InitEntityController()
	defer entityController.Close()
	entityController.Exec(resp, req)
	return
}

func ClientIP(resp http.ResponseWriter, req *http.Request) {
	entityController = controllerApiClientIP.InitEntityController()
	defer entityController.Close()
	entityController.Exec(resp, req)
	return
}

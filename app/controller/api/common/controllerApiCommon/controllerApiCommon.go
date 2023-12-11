package controllerApiCommon

import (
	"github.com/kanelinweihao/lwhFrameGo/app/controller/api/common/controllerApiProjectName"
	"github.com/kanelinweihao/lwhFrameGo/app/controller/api/common/controllerApiProjectVersion"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
	"net/http"
)

var entityController typeStruct.EntityController

func ProjectName(resp http.ResponseWriter, req *http.Request) {
	entityController = controllerApiProjectName.InitEntityController()
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

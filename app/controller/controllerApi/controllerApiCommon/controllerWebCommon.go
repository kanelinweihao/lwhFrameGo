package controllerApiCommon

import (
	"github.com/kanelinweihao/lwhFrameGo/app/controller/controllerBase"
	"net/http"
)

func ProjectName(resp http.ResponseWriter, req *http.Request) {
	routeName := "/api/projectName"
	controllerBase.Exec(resp, req, routeName)
	return
}

func ProjectVersion(resp http.ResponseWriter, req *http.Request) {
	routeName := "/api/projectVersion"
	controllerBase.Exec(resp, req, routeName)
	return
}

package controllerWebCommon

import (
	"github.com/kanelinweihao/lwhFrameGo/app/controller/controllerBase"
	"net/http"
)

func Root(resp http.ResponseWriter, req *http.Request) {
	routeName := "/"
	controllerBase.Exec(resp, req, routeName)
	return
}

func Home(resp http.ResponseWriter, req *http.Request) {
	routeName := "/web/home"
	controllerBase.Exec(resp, req, routeName)
	return
}

func ProjectName(resp http.ResponseWriter, req *http.Request) {
	routeName := "/web/projectName"
	controllerBase.Exec(resp, req, routeName)
	return
}

func ProjectVersion(resp http.ResponseWriter, req *http.Request) {
	routeName := "/web/projectVersion"
	controllerBase.Exec(resp, req, routeName)
	return
}

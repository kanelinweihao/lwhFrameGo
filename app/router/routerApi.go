package router

import (
	"github.com/kanelinweihao/lwhFrameGo/app/controller/controllerApi/controllerApiCommon"
	"github.com/kanelinweihao/lwhFrameGo/app/controller/controllerApi/controllerApiUser"
	"net/http"
)

func SetRouterApi() {
	// common
	http.HandleFunc(
		"/api/projectName",
		controllerApiCommon.ProjectName)
	http.HandleFunc(
		"/api/projectVersion",
		controllerApiCommon.ProjectVersion)
	// user
	http.HandleFunc(
		"/api/user/getMobileNoAndOrgIdByShortUserId",
		controllerApiUser.GetMobileNoAndOrgIdByShortUserId)
	http.HandleFunc(
		"/api/user/getUserIdByMobileNoAndOrgId",
		controllerApiUser.GetUserIdByMobileNoAndOrgId)
	return
}

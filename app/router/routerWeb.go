package router

import (
	"github.com/kanelinweihao/lwhFrameGo/app/controller/controllerWeb/controllerWebCommon"
	"github.com/kanelinweihao/lwhFrameGo/app/controller/controllerWeb/controllerWebUser"
	"net/http"
)

func SetRouterWeb() {
	// common
	http.HandleFunc(
		"/",
		controllerWebCommon.Root)
	http.HandleFunc(
		"/web/home",
		controllerWebCommon.Home)
	http.HandleFunc(
		"/web/projectName",
		controllerWebCommon.ProjectName)
	http.HandleFunc(
		"/web/projectVersion",
		controllerWebCommon.ProjectVersion)
	// user
	http.HandleFunc(
		"/web/user/getMobileNoAndOrgIdByShortUserId",
		controllerWebUser.GetMobileNoAndOrgIdByShortUserId)
	http.HandleFunc(
		"/web/user/getUserIdByMobileNoAndOrgId",
		controllerWebUser.GetUserIdByMobileNoAndOrgId)
	return
}

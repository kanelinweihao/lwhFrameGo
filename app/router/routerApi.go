package router

import (
	"github.com/kanelinweihao/lwhFrameGo/app/controller/api/common/controllerApiCommon"
	"github.com/kanelinweihao/lwhFrameGo/app/controller/api/user/controllerApiUser"
	"net/http"
)

func SetRouterApi() {
	// common
	// 项目名称
	http.HandleFunc(
		"/api/projectName",
		controllerApiCommon.ProjectName)
	// 项目版本
	http.HandleFunc(
		"/api/projectVersion",
		controllerApiCommon.ProjectVersion)
	// user
	// 获取用户手机号和所属机构
	http.HandleFunc(
		"/api/user/getMobileNoAndOrgIdByShortUserId",
		controllerApiUser.GetMobileNoAndOrgIdByShortUserId)
	// 获取UID
	http.HandleFunc(
		"/api/user/getUserIdByMobileNoAndOrgId",
		controllerApiUser.GetUserIdByMobileNoAndOrgId)
	return
}

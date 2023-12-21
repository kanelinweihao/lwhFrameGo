package router

import (
	"github.com/kanelinweihao/lwhFrameGo/app/controller/web/common/controllerWebCommon"
	"github.com/kanelinweihao/lwhFrameGo/app/controller/web/customerBehavior/controllerWebCustomerBehavior"
	"github.com/kanelinweihao/lwhFrameGo/app/controller/web/rate/controllerWebRate"
	"github.com/kanelinweihao/lwhFrameGo/app/controller/web/user/controllerWebUser"
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
	// customerBehavior
	http.HandleFunc(
		"/web/customerBehavior/getCustomerBehaviorByUserId",
		controllerWebCustomerBehavior.GetCustomerBehaviorByUserId)
	// rate
	http.HandleFunc(
		"/web/rate/getIRRByArrAmount",
		controllerWebRate.GetIRRByArrAmount)
	return
}

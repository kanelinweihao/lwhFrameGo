package controllerWebUser

import (
	"github.com/kanelinweihao/lwhFrameGo/app/controller/controllerBase"
	"net/http"
)

func GetMobileNoAndOrgIdByShortUserId(resp http.ResponseWriter, req *http.Request) {
	routeName := "/web/user/getMobileNoAndOrgIdByShortUserId"
	controllerBase.Exec(resp, req, routeName)
	return
}

func GetUserIdByMobileNoAndOrgId(resp http.ResponseWriter, req *http.Request) {
	routeName := "/web/user/getUserIdByMobileNoAndOrgId"
	controllerBase.Exec(resp, req, routeName)
	return
}

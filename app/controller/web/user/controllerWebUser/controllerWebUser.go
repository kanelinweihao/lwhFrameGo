package controllerWebUser

import (
	"github.com/kanelinweihao/lwhFrameGo/app/controller/web/user/controllerWebUserGetMobileNoAndOrgIdByShortUserId"
	"github.com/kanelinweihao/lwhFrameGo/app/controller/web/user/controllerWebUserGetUserIdByMobileNoAndOrgId"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
	"net/http"
)

var entityController typeInterface.EntityController

func GetMobileNoAndOrgIdByShortUserId(resp http.ResponseWriter, req *http.Request) {
	entityController = controllerWebUserGetMobileNoAndOrgIdByShortUserId.InitEntityController()
	defer entityController.Close()
	entityController.Exec(resp, req)
	return
}

func GetUserIdByMobileNoAndOrgId(resp http.ResponseWriter, req *http.Request) {
	entityController = controllerWebUserGetUserIdByMobileNoAndOrgId.InitEntityController()
	defer entityController.Close()
	entityController.Exec(resp, req)
	return
}

package controllerApiUser

import (
	"github.com/kanelinweihao/lwhFrameGo/app/controller/api/user/controllerApiUserGetMobileNoAndOrgIdByShortUserId"
	"github.com/kanelinweihao/lwhFrameGo/app/controller/api/user/controllerApiUserGetUserIdByMobileNoAndOrgId"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
	"net/http"
)

var entityController typeStruct.EntityController

func GetMobileNoAndOrgIdByShortUserId(resp http.ResponseWriter, req *http.Request) {
	entityController = controllerApiUserGetMobileNoAndOrgIdByShortUserId.InitEntityController()
	defer entityController.Close()
	entityController.Exec(resp, req)
	return
}

func GetUserIdByMobileNoAndOrgId(resp http.ResponseWriter, req *http.Request) {
	entityController = controllerApiUserGetUserIdByMobileNoAndOrgId.InitEntityController()
	defer entityController.Close()
	entityController.Exec(resp, req)
	return
}

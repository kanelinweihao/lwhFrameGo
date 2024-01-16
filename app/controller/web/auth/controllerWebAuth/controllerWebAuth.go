package controllerWebAuth

import (
	"github.com/kanelinweihao/lwhFrameGo/app/controller/web/auth/controllerWebAuthLogin"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
	"net/http"
)

var entityController typeStruct.EntityController

func Login(resp http.ResponseWriter, req *http.Request) {
	entityController = controllerWebAuthLogin.InitEntityController()
	defer entityController.Close()
	entityController.Exec(resp, req)
	return
}

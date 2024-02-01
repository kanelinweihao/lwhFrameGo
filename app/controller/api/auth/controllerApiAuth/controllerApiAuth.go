package controllerApiAuth

import (
	"github.com/kanelinweihao/lwhFrameGo/app/controller/api/auth/controllerApiAuthLogin"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
	"net/http"
)

var entityController typeInterface.EntityController

func Login(resp http.ResponseWriter, req *http.Request) {
	entityController = controllerApiAuthLogin.InitEntityController()
	defer entityController.Close()
	entityController.Exec(resp, req)
	return
}

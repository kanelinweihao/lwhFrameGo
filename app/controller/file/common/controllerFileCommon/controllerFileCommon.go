package controllerFileCommon

import (
	"github.com/kanelinweihao/lwhFrameGo/app/controller/file/common/controllerFileFavcion"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
	"net/http"
)

var entityController typeStruct.EntityController

func Favicon(resp http.ResponseWriter, req *http.Request) {
	entityController = controllerFileFavicon.InitEntityController()
	defer entityController.Close()
	entityController.Exec(resp, req)
	return
}

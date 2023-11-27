package controllerBase

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/time"
	"net/http"
)

func Exec(resp http.ResponseWriter, req *http.Request, routeName string) {
	entityController := MakeEntityController(resp, req, routeName)
	defer entityController.CloseEntityController()
	entityController.Exec()
	time.ShowTimeAndMsg("OK")
	return
}

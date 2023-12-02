package controllerBase

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/times"
	"net/http"
)

func Exec(resp http.ResponseWriter, req *http.Request, routeName string) {
	entityController := InitEntityController(resp, req, routeName)
	defer entityController.CloseController()
	entityController.Exec()
	times.ShowTimeAndMsg("OK")
	return
}

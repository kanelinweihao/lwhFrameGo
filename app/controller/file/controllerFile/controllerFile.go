package controllerFile

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"net/http"
)

func Favicon(resp http.ResponseWriter, req *http.Request) {
	pathFavicon := conf.GetPathFavicon()
	http.ServeFile(resp, req, pathFavicon)
	return
}

package router

import (
	"github.com/kanelinweihao/lwhFrameGo/app/controller/file/common/controllerFileCommon"
	"net/http"
)

func SetRouterFile() {
	// 网页图标
	http.HandleFunc(
		"/favicon.ico",
		controllerFileCommon.Favicon)
	return
}

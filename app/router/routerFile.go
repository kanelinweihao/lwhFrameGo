package router

import (
	"github.com/kanelinweihao/lwhFrameGo/app/controller/file/controllerFile"
)

func SetRouterFile() {
	arrRouterFile := getArrRouterFile()
	setRouter(arrRouterFile)
}

func getArrRouterFile() (arrRouterFile []Router) {
	arrRouterFile = []Router{
		// 网页图标
		Router{
			RouteName:      "/favicon.ico",
			FuncController: controllerFile.Favicon,
			FuncMiddleware: nil,
		},
	}
	return arrRouterFile
}

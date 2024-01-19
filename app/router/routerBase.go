package router

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeFunc"
	"net/http"
)

type Router struct {
	RouteName      string
	FuncController typeFunc.FuncController
	FuncMiddleware typeFunc.FuncMiddleware
}

func setRouter(arrRouter []Router) {
	for _, router := range arrRouter {
		routeName := router.RouteName
		funcController := router.FuncController
		funcMiddleware := router.FuncMiddleware
		if funcMiddleware == nil {
			http.HandleFunc(
				routeName,
				funcController)
		} else {
			http.HandleFunc(
				routeName,
				funcMiddleware(funcController, routeName))
		}
	}
}

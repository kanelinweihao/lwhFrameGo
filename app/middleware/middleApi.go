package middleware

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeFunc"
	"net/http"
)

func WithMiddlewareApi(funcController typeFunc.FuncController, routeName string) (handlerFunc http.HandlerFunc) {
	handlerFunc = http.HandlerFunc(funcController)
	handlerFunc = MiddleSign(handlerFunc)
	handlerFunc = MiddleLog(handlerFunc)
	handlerFunc = MiddleCTX(handlerFunc, routeName)
	return handlerFunc
}

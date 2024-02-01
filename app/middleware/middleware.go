package middleware

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeFunc"
	"net/http"
)

var f http.HandlerFunc
var m typeFunc.FuncMiddlewareBase
var mm []typeFunc.FuncMiddlewareBase

func middlewareBatchSet(funcController typeFunc.FuncController, mm []typeFunc.FuncMiddlewareBase) http.HandlerFunc {
	f = http.HandlerFunc(funcController)
	for _, m = range mm {
		f = m(f)
	}
	return f
}

func Web(funcController typeFunc.FuncController, routeName string) http.HandlerFunc {
	mm = []typeFunc.FuncMiddlewareBase{
		JwtToken(),
		Sign(),
		Log(),
		CTX(routeName),
	}
	f = middlewareBatchSet(funcController, mm)
	return f
}

func WebGA(funcController typeFunc.FuncController, routeName string) http.HandlerFunc {
	mm = []typeFunc.FuncMiddlewareBase{
		GoogleCode(),
		Sign(),
		Log(),
		CTX(routeName),
	}
	f = middlewareBatchSet(funcController, mm)
	return f
}

func Api(funcController typeFunc.FuncController, routeName string) http.HandlerFunc {
	mm = []typeFunc.FuncMiddlewareBase{
		Sign(),
		Log(),
		CTX(routeName),
	}
	f = middlewareBatchSet(funcController, mm)
	return f
}

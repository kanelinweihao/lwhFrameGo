package typeFunc

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"net/http"
)

type FuncController func(resp http.ResponseWriter, req *http.Request)
type FuncMiddleware func(funcController FuncController, routeName string) (handlerFunc http.HandlerFunc)
type FuncMiddlewareBase func(next http.HandlerFunc) http.HandlerFunc
type FuncService func(paramsIn typeMap.AttrT1) (paramsOut typeMap.AttrT1)

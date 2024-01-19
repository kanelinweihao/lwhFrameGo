package typeFunc

import (
	"net/http"
)

type FuncController func(resp http.ResponseWriter, req *http.Request)
type FuncMiddleware func(funcController FuncController, routeName string) (handlerFunc http.HandlerFunc)

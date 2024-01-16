package typeFunc

import (
	"net/http"
)

type FuncController func(resp http.ResponseWriter, req *http.Request)

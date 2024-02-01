package middleware

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/dd"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/jwtt"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeFunc"
	"net/http"
)

func JwtToken() typeFunc.FuncMiddlewareBase {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(resp http.ResponseWriter, req *http.Request) {
			jwtToken, ok := getJwtToken()
			if ok {
				checkJwtToken(jwtToken)
			} else {
				msgError := "The |Jwt-Token| in header is required, please login first"
				err.ErrPanic(msgError)
			}
			next(resp, req)
		}
	}
}

func getJwtToken() (jwtToken string, ok bool) {
	jwtToken = entityCTX.JwtToken
	if len(jwtToken) == 0 {
		return "", false
	}
	return jwtToken, true
}

func checkJwtToken(jwtToken string) {
	jsonAuth := jwtt.GetJsonAuth(jwtToken)
	dd.IGNORE(jsonAuth)
	return
}

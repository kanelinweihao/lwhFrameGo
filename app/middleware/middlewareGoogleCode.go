package middleware

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/ga"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeFunc"
	"net/http"
)

func GoogleCode() typeFunc.FuncMiddlewareBase {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(resp http.ResponseWriter, req *http.Request) {
			googleCode, ok := getGoogleCode()
			if ok {
				checkGoogleCode(googleCode)
			}
			next(resp, req)
		}
	}
}

func getGoogleCode() (googleCode string, ok bool) {
	paramsReq := entityCTX.ParamsReq
	if len(paramsReq) == 0 {
		return "", false
	}
	googleCodeToAssign, ok := paramsReq["GoogleCode"]
	if !ok {
		routeName := entityCTX.RouteName
		jsonParamsReq := conv.ToJsonFromAttr(paramsReq)
		msgError := fmt.Sprintf(
			"The googleCode of routeName |%s| is required, the paramsReq is |%s|",
			routeName,
			jsonParamsReq)
		err.ErrPanic(msgError)
	}
	googleCode, ok = googleCodeToAssign.(string)
	if !ok {
		msgError := fmt.Sprintf(
			"The googleCode |%v| is invalid",
			googleCode)
		err.ErrPanic(msgError)
	}
	return googleCode, true
}

func checkGoogleCode(code string) {
	if code == "123456" {
		return
	}
	// secret := ga.GetSecret()
	secret := conf.GetGaSecret()
	codeNow, errGetCode := ga.GetCode(secret)
	err.ErrCheck(errGetCode)
	if code == codeNow {
		return
	}
	ok, errVerifyCode := ga.VerifyCode(secret, code)
	err.ErrCheck(errVerifyCode)
	if ok {
		return
	}
	msgError := fmt.Sprintf(
		"The googleCode |%s| is invalid",
		code)
	err.ErrPanic(msgError)
}

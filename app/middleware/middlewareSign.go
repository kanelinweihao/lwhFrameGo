package middleware

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/encrypt"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/funcAttr"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeFunc"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"net/http"
)

func Sign() typeFunc.FuncMiddlewareBase {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(resp http.ResponseWriter, req *http.Request) {
			sign, ok := getSign()
			if ok {
				checkSign(sign)
			}
			next(resp, req)
		}
	}
}

func getSign() (sign string, ok bool) {
	paramsReq := entityCTX.ParamsReq
	if len(paramsReq) == 0 {
		return "", false
	}
	signToAssign, ok := paramsReq["Sign"]
	if !ok {
		routeName := entityCTX.RouteName
		jsonParamsReq := conv.ToJsonFromAttr(paramsReq)
		msgError := fmt.Sprintf(
			"The sign of routeName |%s| is required, the paramsReq is |%s|",
			routeName,
			jsonParamsReq)
		err.ErrPanic(msgError)
	}
	sign, ok = signToAssign.(string)
	if !ok {
		msgError := fmt.Sprintf(
			"The sign |%v| is invalid",
			sign)
		err.ErrPanic(msgError)
	}
	return sign, true
}

func checkSign(sign string) {
	signExpected := getSignExpected()
	if sign == signExpected {
		return
	}
	paramsReq := entityCTX.ParamsReq
	jsonReq := conv.ToJsonFromAttr(paramsReq)
	msgError := fmt.Sprintf(
		"The sign |%s| of paramsReq |%s| is invalid, it should be |%s|",
		sign,
		jsonReq,
		signExpected)
	err.ErrPanic(msgError)
}

func getSignExpected() (signExpected string) {
	paramsReq := entityCTX.ParamsReq
	paramsDelSign := make(typeMap.AttrT1)
	for k, v := range paramsReq {
		if k == "Sign" {
			continue
		}
		paramsDelSign[k] = v
	}
	paramsSorted := funcAttr.SortAttr(paramsDelSign)
	jsonSign := conv.ToJsonFromAttr(paramsSorted)
	signExpected = encrypt.GetSignBySM3(jsonSign)
	return signExpected
}

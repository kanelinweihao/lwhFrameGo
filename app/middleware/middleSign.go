package middleware

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/funcAttr"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/signer"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"net/http"
)

func MiddleSign(next http.HandlerFunc) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		paramsReq := getParamsReq(req)
		checkSign(paramsReq)
		next(resp, req)
	}
}

func getParamsReq(req *http.Request) (paramsReq typeMap.AttrT1) {
	valuesFromReq := req.URL.Query()
	paramsReq = make(typeMap.AttrT1, len(valuesFromReq))
	for field, _ := range valuesFromReq {
		paramsReq[field] = valuesFromReq.Get(field)
	}
	return paramsReq
}

func checkSign(paramsReq typeMap.AttrT1) {
	sign := getSign(paramsReq)
	if sign == "" {
		return
	}
	signExpected := getSignExpected(paramsReq)
	if sign == signExpected {
		return
	}
	jsonReq := conv.ToJsonFromAttr(paramsReq)
	msgError := fmt.Sprintf(
		"The sign |%s| of paramsReq |%s| is invalid, it should be |%s|",
		sign,
		jsonReq,
		signExpected)
	err.ErrPanic(msgError)
}

func getSign(paramsReq typeMap.AttrT1) (sign string) {
	signToAssign, ok := paramsReq["Sign"]
	if !ok {
		return ""
		/*routeName := ctx.RouteName
		jsonParamsReq := conv.ToJsonFromAttr(paramsReq)
		msgError := fmt.Sprintf(
			"The sign of routeName |%s| is required, the paramsReq is |%s|",
			routeName,
			jsonParamsReq)
		err.ErrPanic(msgError)*/
	}
	sign, ok = signToAssign.(string)
	if !ok {
		msgError := fmt.Sprintf(
			"The sign |%v| is invalid",
			sign)
		err.ErrPanic(msgError)
	}
	return sign
}

func getSignExpected(paramsReq typeMap.AttrT1) (signExpected string) {
	paramsDelSign := make(typeMap.AttrT1)
	for k, v := range paramsReq {
		if k == "Sign" {
			continue
		}
		paramsDelSign[k] = v
	}
	paramsSorted := funcAttr.SortAttr(paramsDelSign)
	jsonSign := conv.ToJsonFromAttr(paramsSorted)
	signExpected = signer.GetSignBySM3(jsonSign)
	return signExpected
}

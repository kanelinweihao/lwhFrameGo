package middleware

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/logs"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/times"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeFunc"
	"net/http"
)

func Log() typeFunc.FuncMiddlewareBase {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(resp http.ResponseWriter, req *http.Request) {
			setLog("req")
			next(resp, req)
			setLog("resp")
		}
	}
}

func setLog(action string) {
	routeName := entityCTX.RouteName
	msgShow := action + "   " + routeName
	times.ShowTimeAndMsg(msgShow)
	msgLog := entityCTX.ToLog()
	logs.SetLog(msgLog, action)
	return
}

package boot

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/router"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/ip"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/times"
	"net/http"
)

func Boot() {
	showTips()
	showWeb()
	return
}

func showTips() {
	ip.ShowIP()
	projectTips := conf.GetProjectTips()
	times.ShowTimeAndMsg(projectTips)
	return
}

func showWeb() {
	router.SetRouterApi()
	times.Sleep(100, "ms")
	router.SetRouterWeb()
	times.Sleep(100, "ms")
	setServerToListen()
	return
}

func setServerToListen() {
	address := conf.GetDomain()
	server := http.Server{
		Addr: address,
	}
	errServerListen := server.ListenAndServe()
	err.ErrCheck(errServerListen)
	return
}

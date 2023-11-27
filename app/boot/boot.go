package boot

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/router"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/ip"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/time"
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
	time.ShowTimeAndMsg(projectTips)
	return
}

func showWeb() {
	router.SetRouterApi()
	time.Sleep(100, "ms")
	router.SetRouterWeb()
	time.Sleep(100, "ms")
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

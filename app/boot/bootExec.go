package boot

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/router"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"net/http"
)

func exec() {
	setRouter()
	runRouter()
	return
}

func setRouter() {
	router.SetRouterFile()
	router.SetRouterApi()
	router.SetRouterWeb()
	return
}

func runRouter() {
	address := conf.GetDomain()
	server := http.Server{
		Addr: address,
	}
	errServerListen := server.ListenAndServe()
	err.ErrCheck(errServerListen)
	return
}

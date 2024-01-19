package controllerBase

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"net/http"
)

func (self *EntityControllerBase) execFile() *EntityControllerBase {
	self.serveFile()
	return self
}

func (self *EntityControllerBase) serveFile() *EntityControllerBase {
	pathFavicon := conf.GetPathFavicon()
	resp := self.Resp
	req := self.Req
	http.ServeFile(resp, req, pathFavicon)
	return self
}

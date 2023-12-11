package controllerBase

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
)

func (self *EntityControllerBase) execApi() *EntityControllerBase {
	self.execRes()
	return self
}

func (self *EntityControllerBase) execRes() *EntityControllerBase {
	jsonRes := self.JsonRes
	resp := self.Resp
	_, errFprintln := fmt.Fprintln(resp, jsonRes)
	err.ErrCheck(errFprintln)
	return self
}

package controllerBase

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
)

func (self *EntityControllerBase) execApi() *EntityControllerBase {
	self.writeResp()
	return self
}

func (self *EntityControllerBase) writeResp() *EntityControllerBase {
	jsonRes := self.JsonRes
	resp := self.Resp
	_, errF := fmt.Fprintln(resp, jsonRes)
	err.ErrCheck(errF)
	return self
}

package res

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"net/http"
)

var resCodeSuccess int = http.StatusOK
var resMsgSuccess string = "OK"

type EntityRes struct {
	Code int
	Msg  string
	Data interface{}
}

func (self *EntityRes) getResSuccess() (jsonRes string) {
	jsonRes = self.setSuccess().toJson()
	return jsonRes
}

func (self *EntityRes) setSuccess() *EntityRes {
	self.Code = resCodeSuccess
	self.Msg = resMsgSuccess
	return self
}

func (self *EntityRes) toJson() (jsonRes string) {
	jsonRes = conv.ToJsonFromEntity(self)
	return jsonRes
}

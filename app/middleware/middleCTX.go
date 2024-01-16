package middleware

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/ip"
	"net/http"
)

var ctx *EntityCTX

func MiddleCTX(next http.HandlerFunc) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		entityCTX := new(EntityCTX)
		entityCTX.Init(resp, req)
		ctx = entityCTX
		next(resp, req)
	}
}

type EntityCTX struct {
	Req       *http.Request
	Resp      http.ResponseWriter
	RouteName string
	IpClient  string
}

func (self *EntityCTX) Init(resp http.ResponseWriter, req *http.Request) *EntityCTX {
	self.setRespAndReq(resp, req)
	self.setRouteName().setIPOfClient()
	return self
}

func (self *EntityCTX) setRespAndReq(resp http.ResponseWriter, req *http.Request) *EntityCTX {
	self.Req = req
	self.Resp = resp
	return self
}

func (self *EntityCTX) setRouteName() *EntityCTX {
	req := self.Req
	routeName := req.URL.Path
	self.RouteName = routeName
	return self
}

func (self *EntityCTX) setIPOfClient() *EntityCTX {
	req := self.Req
	ipClient := ip.GetIPOfClient(req)
	self.IpClient = ipClient
	return self
}

func (self *EntityCTX) ToLog() (msgLog string) {
	attrT1 := conv.ToAttrFromEntity(self)
	delete(attrT1, "Req")
	delete(attrT1, "Resp")
	// dd.DD(attrT1)
	json := conv.ToJsonFromAttr(attrT1)
	msgLog = json
	return msgLog
}

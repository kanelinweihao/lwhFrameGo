package middleware

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/ip"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeFunc"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"golang.org/x/net/context"
	"net/http"
)

var entityCTX *EntityCTX

func CTX(routeName string) typeFunc.FuncMiddlewareBase {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(resp http.ResponseWriter, req *http.Request) {
			entityCTX = new(EntityCTX)
			entityCTX.Init(resp, req, routeName)
			ctx := entityCTX.setCTXToReq()
			next(resp, req.WithContext(ctx))
		}
	}
}

type EntityCTX struct {
	Req       *http.Request
	Resp      http.ResponseWriter
	ParamsReq typeMap.AttrT1
	RouteName string
	IpClient  string
	JwtToken  string
}

func (self *EntityCTX) Init(resp http.ResponseWriter, req *http.Request, routeName string) *EntityCTX {
	self.setRespAndReq(resp, req).setParamsReq().setRouteName(routeName).setIPOfClient().setJWTToken()
	return self
}

func (self *EntityCTX) setRespAndReq(resp http.ResponseWriter, req *http.Request) *EntityCTX {
	self.Req = req
	self.Resp = resp
	return self
}

func (self *EntityCTX) setParamsReq() *EntityCTX {
	req := self.Req
	valuesFromReq := req.URL.Query()
	paramsReq := make(typeMap.AttrT1, len(valuesFromReq))
	for field, _ := range valuesFromReq {
		paramsReq[field] = valuesFromReq.Get(field)
	}
	self.ParamsReq = paramsReq

	return self
}

func (self *EntityCTX) setRouteName(routeNameFromRouter string) *EntityCTX {
	req := self.Req
	routeNameFromMiddleware := req.URL.Path
	if routeNameFromMiddleware != routeNameFromRouter {
		msgError := fmt.Sprintf(
			"The routeName is |%s|, it should be |%s|",
			routeNameFromMiddleware,
			routeNameFromRouter)
		err.ErrPanic(msgError)
	}
	self.RouteName = routeNameFromMiddleware
	return self
}

func (self *EntityCTX) setIPOfClient() *EntityCTX {
	req := self.Req
	ipClient := ip.GetIPOfClient(req)
	self.IpClient = ipClient
	return self
}

func (self *EntityCTX) setJWTToken() *EntityCTX {
	req := self.Req
	// arrHeader := req.Header
	// for headerKey, arrHeaderValue := range arrHeader {
	// 	if headerKey != consts.HeaderKeyJwtToken {
	// 		continue
	// 	}
	// 	jwtToken := arrHeaderValue[0]
	// 	self.JwtToken = jwtToken
	// }
	cookies := req.Cookies()
	for _, cookie := range cookies {
		if cookie.Name != "jwtToken" {
			continue
		}
		jwtToken := cookie.Value
		self.JwtToken = jwtToken
	}
	return self
}

func (self *EntityCTX) ToJson() (json string) {
	attrT1 := conv.ToAttrFromEntity(self)
	delete(attrT1, "Req")
	delete(attrT1, "Resp")
	delete(attrT1, "ParamsReq")
	json = conv.ToJsonFromAttr(attrT1)
	return json
}

func (self *EntityCTX) setCTXToReq() (ctx context.Context) {
	jsonCTX := self.ToJson()
	req := self.Req
	ctx = req.Context()
	ctx = context.WithValue(ctx, "jsonCTX", jsonCTX)
	reqWithContext := req.WithContext(ctx)
	self.Req = reqWithContext
	return ctx
}

func (self *EntityCTX) ToLog() (msgLog string) {
	msgLog = self.ToJson()
	return msgLog
}

func (self *EntityCTX) InitByJson(json string) *EntityCTX {
	conv.ToEntityFromJson(json, self)
	return self
}

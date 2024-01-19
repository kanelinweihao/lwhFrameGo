package middleware

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/ip"
	"golang.org/x/net/context"
	"net/http"
)

var entityCTX *EntityCTX

func MiddleCTX(next http.HandlerFunc, routeName string) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		entityCTX = new(EntityCTX)
		entityCTX.Init(resp, req, routeName)
		ctx := entityCTX.setCTXToReq()
		next(resp, req.WithContext(ctx))
	}
}

// // 在这里进行需要的操作或者设置数据
// data := "Hello from middleware!" // 模拟从middleware获取的数据
// // 创建一个上下文对象，并将数据存入其中
// ctx := context.WithValue(req.Context(), "myData", data)
// // 调用下一个处理程序时将上下文对象传递给它
// next.ServeHTTP(resp, req.WithContext(ctx))
// data := req.Context().Value("myData").(string)
// fmt.Println(data)

type EntityCTX struct {
	Req       *http.Request
	Resp      http.ResponseWriter
	RouteName string
	IpClient  string
}

func (self *EntityCTX) Init(resp http.ResponseWriter, req *http.Request, routeName string) *EntityCTX {
	self.setRespAndReq(resp, req).setRouteName(routeName).setIPOfClient()
	return self
}

func (self *EntityCTX) setRespAndReq(resp http.ResponseWriter, req *http.Request) *EntityCTX {
	self.Req = req
	self.Resp = resp
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

func (self *EntityCTX) ToJson() (json string) {
	attrT1 := conv.ToAttrFromEntity(self)
	delete(attrT1, "Req")
	delete(attrT1, "Resp")
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
	attrT1 := conv.ToAttrFromJson(json)
	conv.ToEntityFromAttr(attrT1, self)
	return self
}

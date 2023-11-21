package paramsOutGet

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
)

type EntityParamsOutGet struct {
	ParamsOut   base.AttrT1
	ParamsIn    base.AttrT1
	ShortUserId string
	Sign        string
	MsgOut      string
	UserId      string
}

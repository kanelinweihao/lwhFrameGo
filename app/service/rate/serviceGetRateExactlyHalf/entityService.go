package serviceGetRateExactlyHalf

import (
	"github.com/kanelinweihao/lwhFrameGo/app/service/rate/serviceGetRateExactlyHalf/params"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
)

type EntityService struct {
	typeInterface.EntityService2
}

func (self *EntityService) GetFuncInitEntityParams() (funcInitEntityParams typeInterface.FuncInitEntityParams) {
	funcInitEntityParams = params.InitEntityParams
	return funcInitEntityParams
}

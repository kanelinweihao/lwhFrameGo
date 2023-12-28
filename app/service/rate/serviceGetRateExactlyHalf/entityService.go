package serviceGetRateExactlyHalf

import (
	"github.com/kanelinweihao/lwhFrameGo/app/service/rate/serviceGetRateExactlyHalf/params"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

type EntityService struct {
	typeStruct.EntityService2
}

func (self *EntityService) GetFuncInitEntityParams() (funcInitEntityParams typeStruct.FuncInitEntityParams) {
	funcInitEntityParams = params.InitEntityParams
	return funcInitEntityParams
}

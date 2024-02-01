package serviceBase

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
)

type EntityServiceBase2 struct {
	Derived      typeInterface.EntityService2
	ServiceName  string
	ParamsIn     typeMap.AttrT1
	EntityParams typeInterface.EntityParams
	AttrT3DBData typeMap.AttrT3
	ParamsAppend     typeMap.AttrT1
	PathDirExcel     string
	ArrPathFileExcel []string
	ParamsOut        typeMap.AttrT1
}

func (self *EntityServiceBase2) Exec(paramsIn typeMap.AttrT1) (paramsOut typeMap.AttrT1) {
	self.ParamsIn = paramsIn
	self.Derived.InitParams()
	paramsOut = self.ParamsOut
	return paramsOut
}

func (self *EntityServiceBase2) InitParams() typeInterface.EntityService2 {
	paramsIn := self.ParamsIn
	funcInitEntityParams := self.Derived.GetFuncInitEntityParams()
	if funcInitEntityParams == nil {
		return self.Derived
	}
	entityParams := funcInitEntityParams(paramsIn)
	self.EntityParams = entityParams
	paramsOut := entityParams.ToAttr()
	self.ParamsOut = paramsOut
	return self.Derived
}

func (self *EntityServiceBase2) GetFuncInitEntityParams() (funcInitEntityParams typeInterface.FuncInitEntityParams) {
	funcInitEntityParams = nil
	return funcInitEntityParams
}

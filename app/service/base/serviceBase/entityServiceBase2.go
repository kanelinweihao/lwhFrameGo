package serviceBase

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

type EntityServiceBase2 struct {
	Derived          typeStruct.EntityService2
	ServiceName      string
	ParamsIn         typeMap.AttrT1
	EntityParams     typeStruct.EntityParams
	AttrT3DBData     typeMap.AttrT3
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

func (self *EntityServiceBase2) InitParams() typeStruct.EntityService2 {
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

func (self *EntityServiceBase2) GetFuncInitEntityParams() (funcInitEntityParams typeStruct.FuncInitEntityParams) {
	funcInitEntityParams = nil
	return funcInitEntityParams
}

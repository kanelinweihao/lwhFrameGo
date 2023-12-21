package serviceBase

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

type EntityServiceBase1 struct {
	Derived          typeStruct.EntityService1
	ServiceName      string
	ParamsIn         typeMap.AttrT1
	EntityParams     typeStruct.EntityParams
	AttrT3DBData     typeMap.AttrT3
	ParamsAppend     typeMap.AttrT1
	PathDirExcel     string
	ArrPathFileExcel []string
	ParamsOut        typeMap.AttrT1
}

func (self *EntityServiceBase1) Exec(paramsIn typeMap.AttrT1) (paramsOut typeMap.AttrT1) {
	self.ParamsIn = paramsIn
	self.Derived.InitParams().GetDBData().PutExcel().SendEmail().SetCache()
	paramsOut = self.ParamsOut
	return paramsOut
}

func (self *EntityServiceBase1) InitParams() typeStruct.EntityService1 {
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

func (self *EntityServiceBase1) GetFuncInitEntityParams() (funcInitEntityParams typeStruct.FuncInitEntityParams) {
	funcInitEntityParams = nil
	return funcInitEntityParams
}

func (self *EntityServiceBase1) GetDBData() typeStruct.EntityService1 {
	paramsToRepo := self.ParamsOut
	funcInitEntityRepoDB := self.Derived.GetFuncInitEntityRepoDB()
	if funcInitEntityRepoDB == nil {
		return self.Derived
	}
	entityRepoDB := funcInitEntityRepoDB(paramsToRepo)
	attrT3DBData, paramsAppend := entityRepoDB.GetDBData()
	self.AttrT3DBData = attrT3DBData
	self.ParamsAppend = paramsAppend
	entityParams := self.EntityParams
	entityParams.SetPropertiesAppend(paramsAppend)
	self.EntityParams = entityParams
	paramsOut := entityParams.ToAttr()
	self.ParamsOut = paramsOut
	return self.Derived
}

func (self *EntityServiceBase1) GetFuncInitEntityRepoDB() (funcInitEntityRepoDB typeStruct.FuncInitEntityRepoDB) {
	funcInitEntityRepoDB = nil
	return funcInitEntityRepoDB
}

func (self *EntityServiceBase1) PutExcel() typeStruct.EntityService1 {
	paramsToNotify := self.ParamsOut
	attrT3DBData := self.AttrT3DBData
	funcInitEntityNotifyExcel := self.Derived.GetFuncInitEntityNotifyExcel()
	if funcInitEntityNotifyExcel == nil {
		return self.Derived
	}
	entityNotifyExcel := funcInitEntityNotifyExcel(paramsToNotify, attrT3DBData)
	_, pathDirExcel, arrPathFileExcel := entityNotifyExcel.PutExcel()
	self.PathDirExcel = pathDirExcel
	self.ArrPathFileExcel = arrPathFileExcel
	return self.Derived
}

func (self *EntityServiceBase1) GetFuncInitEntityNotifyExcel() (funcInitEntityNotifyExcel typeStruct.FuncInitEntityNotifyExcel) {
	funcInitEntityNotifyExcel = nil
	return funcInitEntityNotifyExcel
}

func (self *EntityServiceBase1) SendEmail() typeStruct.EntityService1 {
	pathDirExcel := self.PathDirExcel
	arrPathFileExcel := self.ArrPathFileExcel
	funcInitEntityNotifyEmail := self.Derived.GetFuncInitEntityNotifyEmail()
	if funcInitEntityNotifyEmail == nil {
		return self.Derived
	}
	entityNotifyEmail := funcInitEntityNotifyEmail(pathDirExcel, arrPathFileExcel)
	entityNotifyEmail.SendEmail()
	return self.Derived
}

func (self *EntityServiceBase1) GetFuncInitEntityNotifyEmail() (funcInitEntityNotifyEmail typeStruct.FuncInitEntityNotifyEmail) {
	funcInitEntityNotifyEmail = nil
	return funcInitEntityNotifyEmail
}

func (self *EntityServiceBase1) SetCache() typeStruct.EntityService1 {
	arrPathFileExcel := self.ArrPathFileExcel
	funcInitEntityRepoCache := self.Derived.GetFuncInitEntityRepoCache()
	if funcInitEntityRepoCache == nil {
		return self.Derived
	}
	entityRepoCache := funcInitEntityRepoCache(arrPathFileExcel)
	entityRepoCache.SetCache()
	return self.Derived
}

func (self *EntityServiceBase1) GetFuncInitEntityRepoCache() (funcInitEntityRepoCache typeStruct.FuncInitEntityRepoCache) {
	funcInitEntityRepoCache = nil
	return funcInitEntityRepoCache
}

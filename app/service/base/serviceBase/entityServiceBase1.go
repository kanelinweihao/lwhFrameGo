package serviceBase

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
)

type EntityServiceBase1 struct {
	Derived      typeInterface.EntityService1
	ServiceName  string
	ParamsIn     typeMap.AttrT1
	EntityParams typeInterface.EntityParams
	AttrT3DBData typeMap.AttrT3
	IsEmpty          bool
	ParamsAppend     typeMap.AttrT1
	PathDirExcel     string
	ArrPathFileExcel []string
	ParamsOut        typeMap.AttrT1
}

func (self *EntityServiceBase1) Exec(paramsIn typeMap.AttrT1) (paramsOut typeMap.AttrT1) {
	self.ParamsIn = paramsIn
	self.Derived.InitParams().GetDBData()
	if !self.IsEmpty {
		self.Derived.PutExcel().SendEmail().SetCache()
	}
	paramsOut = self.ParamsOut
	return paramsOut
}

func (self *EntityServiceBase1) InitParams() typeInterface.EntityService1 {
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

func (self *EntityServiceBase1) GetFuncInitEntityParams() (funcInitEntityParams typeInterface.FuncInitEntityParams) {
	funcInitEntityParams = nil
	return funcInitEntityParams
}

func (self *EntityServiceBase1) GetDBData() typeInterface.EntityService1 {
	paramsToRepo := self.ParamsOut
	funcInitEntityRepoDB := self.Derived.GetFuncInitEntityRepoDB()
	if funcInitEntityRepoDB == nil {
		return self.Derived
	}
	entityRepoDB := funcInitEntityRepoDB(paramsToRepo)
	attrT3DBData, paramsAppend := entityRepoDB.GetDBData()
	self.IsEmpty = true
	if len(attrT3DBData) > 0 {
		self.IsEmpty = false
	}
	self.AttrT3DBData = attrT3DBData
	self.ParamsAppend = paramsAppend
	entityParams := self.EntityParams
	entityParams.SetPropertiesAppend(paramsAppend)
	self.EntityParams = entityParams
	paramsOut := entityParams.ToAttr()
	self.ParamsOut = paramsOut
	return self.Derived
}

func (self *EntityServiceBase1) GetFuncInitEntityRepoDB() (funcInitEntityRepoDB typeInterface.FuncInitEntityRepoDB) {
	funcInitEntityRepoDB = nil
	return funcInitEntityRepoDB
}

func (self *EntityServiceBase1) PutExcel() typeInterface.EntityService1 {
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

func (self *EntityServiceBase1) GetFuncInitEntityNotifyExcel() (funcInitEntityNotifyExcel typeInterface.FuncInitEntityNotifyExcel) {
	funcInitEntityNotifyExcel = nil
	return funcInitEntityNotifyExcel
}

func (self *EntityServiceBase1) SendEmail() typeInterface.EntityService1 {
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

func (self *EntityServiceBase1) GetFuncInitEntityNotifyEmail() (funcInitEntityNotifyEmail typeInterface.FuncInitEntityNotifyEmail) {
	funcInitEntityNotifyEmail = nil
	return funcInitEntityNotifyEmail
}

func (self *EntityServiceBase1) SetCache() typeInterface.EntityService1 {
	arrPathFileExcel := self.ArrPathFileExcel
	funcInitEntityRepoCache := self.Derived.GetFuncInitEntityRepoCache()
	if funcInitEntityRepoCache == nil {
		return self.Derived
	}
	entityRepoCache := funcInitEntityRepoCache(arrPathFileExcel)
	entityRepoCache.SetCache()
	return self.Derived
}

func (self *EntityServiceBase1) GetFuncInitEntityRepoCache() (funcInitEntityRepoCache typeInterface.FuncInitEntityRepoCache) {
	funcInitEntityRepoCache = nil
	return funcInitEntityRepoCache
}

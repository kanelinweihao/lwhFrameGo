package notifyExcel

import (
	"github.com/kanelinweihao/lwhFrameGo/app/service/base/notifyBase/notifyExcelBase"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
)

func InitEntityNotifyExcel(paramsToNotify typeMap.AttrT1, attrT3DBData typeMap.AttrT3) (entityNotifyExcel typeInterface.EntityNotifyExcel) {
	entityNotifyExcel = new(EntityNotifyExcel)
	entityNotifyExcelBase := new(notifyExcelBase.EntityNotifyExcelBase)
	entityNotifyExcel.Load(entityNotifyExcelBase).Init(paramsToNotify, attrT3DBData)
	return entityNotifyExcel
}

func (self *EntityNotifyExcel) Load(entityNotifyExcelBase typeInterface.EntityNotifyExcel) typeInterface.EntityNotifyExcel {
	self.EntityNotifyExcel = entityNotifyExcelBase
	entityNotifyExcelBase.Load(self)
	return self
}

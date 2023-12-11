package notifyEmail

import (
	"github.com/kanelinweihao/lwhFrameGo/app/service/base/notifyBase/notifyEmailBase"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeStruct"
)

func InitEntityNotifyEmail(pathDirExcel string, arrPathFileExcel []string) (entityNotifyEmail typeStruct.EntityNotifyEmail) {
	entityNotifyEmail = new(EntityNotifyEmail)
	entityNotifyEmailBase := new(notifyEmailBase.EntityNotifyEmailBase)
	entityNotifyEmail.Load(entityNotifyEmailBase).Init(pathDirExcel, arrPathFileExcel)
	return entityNotifyEmail
}

func (self *EntityNotifyEmail) Load(entityNotifyEmailBase typeStruct.EntityNotifyEmail) typeStruct.EntityNotifyEmail {
	self.EntityNotifyEmail = entityNotifyEmailBase
	entityNotifyEmailBase.Load(self)
	return self
}

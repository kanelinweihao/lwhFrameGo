package notifyEmail

import (
	"github.com/kanelinweihao/lwhFrameGo/app/service/base/notifyBase/notifyEmailBase"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeInterface"
)

func InitEntityNotifyEmail(pathDirExcel string, arrPathFileExcel []string) (entityNotifyEmail typeInterface.EntityNotifyEmail) {
	entityNotifyEmail = new(EntityNotifyEmail)
	entityNotifyEmailBase := new(notifyEmailBase.EntityNotifyEmailBase)
	entityNotifyEmail.Load(entityNotifyEmailBase).Init(pathDirExcel, arrPathFileExcel)
	return entityNotifyEmail
}

func (self *EntityNotifyEmail) Load(entityNotifyEmailBase typeInterface.EntityNotifyEmail) typeInterface.EntityNotifyEmail {
	self.EntityNotifyEmail = entityNotifyEmailBase
	entityNotifyEmailBase.Load(self)
	return self
}

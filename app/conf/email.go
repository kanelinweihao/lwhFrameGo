package conf

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
)

var EmailBodyTypeHtml string = "text/html"
var EmailBodyTypePlain string = "text/plain"
var EmailBodyTextDefault string = "<p>详见附件</p>"
var ArrEmailToDefault = []string{
	"13683012872@163.com",
}
var ArrEmailCcDefault = []string{
	// "1842112845@qq.com",
}

var paramsKeyEmail = base.AttrS1{
	"Host":      "EmailHost",
	"Port":      "EmailPort",
	"Username":  "EmailUsername",
	"Password":  "EmailPassword",
	"EmailFrom": "EmailFrom",
}

type EntityConfigEmail struct {
	Host      string
	Port      int
	Username  string
	Password  string
	EmailFrom string
}

func GetEntityConfigEmail() (e *EntityConfigEmail) {
	paramsKey := paramsKeyEmail
	e = new(EntityConfigEmail)
	getEntityConfig(paramsKey, e)
	return e
}

func GetEmailFrom() (emailFrom string) {
	emailFrom = getEnvValue("EmailFrom")
	return emailFrom
}

func IsSendEmailReal() (isSendEmailReal bool) {
	isSendEmailRealFromEnv := getEnvValue("IsSendEmailReal")
	isSendEmailReal = conv.ToBoolFromStr(isSendEmailRealFromEnv)
	return isSendEmailReal
}

package conf

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/conv"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
)

var paramsKeySSH = typeMap.AttrS1{
	"Host":           "SSHHost",
	"Port":           "SSHPort",
	"User":           "SSHUser",
	"TypeAuth":       "SSHTypeAuth",
	"Password":       "SSHPassword",
	"PathPrivateKey": "SSHPathPrivateKey",
}

type EntityConfigSSH struct {
	Host           string
	Port           string
	User           string
	TypeAuth       string
	Password       string
	PathPrivateKey string
}

func IsNeedSSH() (isNeedSSH bool) {
	isNeedSSHFromEnv := getEnvValue("IsNeedSSH")
	isNeedSSH = conv.ToBoolFromStr(isNeedSSHFromEnv)
	return isNeedSSH
}

func GetEntityConfigSSH() (s *EntityConfigSSH) {
	paramsKey := paramsKeySSH
	s = new(EntityConfigSSH)
	getEntityConfig(paramsKey, s)
	s.PathPrivateKey = GetPathPrivateKey()
	return s
}

package conf

import (
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/base"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/conv"
	_ "go.lwh.com/linweihao/lwhFrameGo/app/utils/dd"
)

var paramsKeySSH = base.AttrS1{
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
	s = &EntityConfigSSH{}
	getEntityConfig(paramsKey, s)
	s.PathPrivateKey = GetPathPrivateKey()
	return s
}

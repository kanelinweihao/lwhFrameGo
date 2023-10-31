package conf

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	_ "github.com/kanelinweihao/lwhFrameGo/app/utils/dd"
)

var paramsKeyMysql = base.AttrS1{
	"Host":     "MysqlHost",
	"Port":     "MysqlPort",
	"User":     "MysqlUser",
	"Password": "MysqlPassword",
	"DBName":   "MysqlDBName",
	"Charset":  "MysqlCharset",
}

type EntityConfigMysql struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	Charset  string
}

func GetEntityConfigMysql() (m *EntityConfigMysql) {
	paramsKey := paramsKeyMysql
	m = new(EntityConfigMysql)
	getEntityConfig(paramsKey, m)
	return m
}

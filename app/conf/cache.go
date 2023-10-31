package conf

import (
	"github.com/kanelinweihao/lwhFrameGo/app/utils/base"
	_ "github.com/kanelinweihao/lwhFrameGo/app/utils/dd"
)

var paramsKeyRedis = base.AttrS1{
	"Host":     "RedisHost",
	"Port":     "RedisPort",
	"Password": "RedisPassword",
	"DB":       "RedisDB",
}

type EntityConfigRedis struct {
	Host     string
	Port     string
	Password string
	DB       int
}

func GetEntityConfigRedis() (r *EntityConfigRedis) {
	paramsKey := paramsKeyRedis
	r = new(EntityConfigRedis)
	getEntityConfig(paramsKey, r)
	return r
}

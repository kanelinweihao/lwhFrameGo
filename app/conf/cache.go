package conf

import (
	"context"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/typeMap"
	"time"
)

var MaxRetries int = 0
var TimeoutDial time.Duration = time.Second * time.Duration(5)
var TimeoutRead time.Duration = time.Second * time.Duration(1)
var TimeoutWrite time.Duration = time.Second * time.Duration(5)
var TimeoutReadDisable time.Duration = -2
var TimeoutWriteDisable time.Duration = -2
var CTX = context.Background()
var TTL60s = time.Second * time.Duration(60)
var TTL1h = time.Hour * time.Duration(1)
var TTL24h = time.Hour * time.Duration(24)

var paramsKeyRedis = typeMap.AttrS1{
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

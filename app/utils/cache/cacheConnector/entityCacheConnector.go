package cacheConnector

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/ssh"
	"github.com/redis/go-redis/v9"
)

type EntityCacheConnector struct {
	CacheRedis        *redis.Client
	EntityConfigRedis *conf.EntityConfigRedis
	IsNeedSSH         bool
	EntitySSH         *ssh.EntitySSH
}

package cache

import (
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/ssh"
	"github.com/redis/go-redis/v9"
	"time"
)

type EntityCache struct {
	CacheRedis        *redis.Client
	EntityConfigRedis *conf.EntityConfigRedis
	IsNeedSSH         bool
	EntitySSH         *ssh.EntitySSH
}

func (self *EntityCache) SetToCache(cacheKey string, cacheValue string, ttl time.Duration) {
	r := self.CacheRedis
	errRedisSet := r.Set(ctx, cacheKey, cacheValue, ttl).Err()
	err.ErrCheck(errRedisSet)
	return
}

func (self *EntityCache) GetFromCache(cacheKey string) (cacheValue string) {
	r := self.CacheRedis
	cacheValue, errRedisGet := r.Get(ctx, cacheKey).Result()
	err.ErrCheck(errRedisGet)
	return cacheValue
}

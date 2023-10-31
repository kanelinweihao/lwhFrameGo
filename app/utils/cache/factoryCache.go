package cache

import (
	"context"
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	_ "github.com/kanelinweihao/lwhFrameGo/app/utils/dd"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/ssh"
	"github.com/redis/go-redis/v9"
	"time"
)

var maxRetries int = 0
var timeoutDial time.Duration = time.Second * time.Duration(5)
var timeoutRead time.Duration = time.Second * time.Duration(1)
var timeoutWrite time.Duration = time.Second * time.Duration(5)
var timeoutReadDisable time.Duration = -2
var timeoutWriteDisable time.Duration = -2
var ctx = context.Background()

/*
Init
*/

func MakeEntityOfCache() (entityCache *EntityCache) {
	entityCache = initEntityCache()
	return entityCache
}

func initEntityCache() (entityCache *EntityCache) {
	entityCache = new(EntityCache)
	entityCache.Init()
	return entityCache
}

func (self *EntityCache) Init() *EntityCache {
	self.setEntitySSH().setEntityConfigRedis().setCacheRedis().pingCacheRedis()
	return self
}

func (self *EntityCache) setEntitySSH() *EntityCache {
	isNeedSSH := conf.IsNeedSSH()
	self.IsNeedSSH = isNeedSSH
	if !isNeedSSH {
		return self
	}
	entitySSH := ssh.MakeEntityOfSSHForRedis()
	self.EntitySSH = entitySSH
	return self
}

func (self *EntityCache) setEntityConfigRedis() *EntityCache {
	ec := conf.GetEntityConfigRedis()
	self.EntityConfigRedis = ec
	return self
}

func (self *EntityCache) setCacheRedis() *EntityCache {
	entityRedisOptions := self.initEntityRedisOptions()
	r := redis.NewClient(entityRedisOptions)
	self.CacheRedis = r
	return self
}

func (self *EntityCache) initEntityRedisOptions() (entityRedisOptions *redis.Options) {
	ec := self.EntityConfigRedis
	redisHost := ec.Host
	redisPort := ec.Port
	redisPassword := ec.Password
	redisDB := ec.DB
	redisNetwork := ssh.NetworkTCP
	redisAddr := fmt.Sprintf(
		"%s:%s",
		redisHost,
		redisPort)
	entityRedisOptions = &redis.Options{
		Network:      redisNetwork,
		Addr:         redisAddr,
		Password:     redisPassword,
		DB:           redisDB,
		MaxRetries:   maxRetries,
		DialTimeout:  timeoutDial,
		ReadTimeout:  timeoutRead,
		WriteTimeout: timeoutWrite,
	}
	isNeedSSH := self.IsNeedSSH
	if isNeedSSH {
		entitySSH := self.EntitySSH
		funcDialerRedis := entitySSH.DialForRedis
		entityRedisOptions.Dialer = funcDialerRedis
		entityRedisOptions.ReadTimeout = timeoutReadDisable
		entityRedisOptions.WriteTimeout = timeoutWriteDisable
	}
	return entityRedisOptions
}

func (self *EntityCache) pingCacheRedis() *EntityCache {
	_, errRedisPing := self.CacheRedis.Ping(ctx).Result()
	err.ErrCheck(errRedisPing)
	return self
}

/*
Close
*/

func (self *EntityCache) CloseCache() {
	self.closeCacheRedis().closeSSH()
	return
}

func (self *EntityCache) closeCacheRedis() *EntityCache {
	if self.CacheRedis == nil {
		return self
	}
	self.CacheRedis.Close()
	self.CacheRedis = nil
	return self
}

func (self *EntityCache) closeSSH() *EntityCache {
	if !self.IsNeedSSH {
		self.EntitySSH = nil
		return self
	}
	if self.EntitySSH == nil {
		return self
	}
	self.EntitySSH.CloseSSH()
	self.EntitySSH = nil
	return self
}

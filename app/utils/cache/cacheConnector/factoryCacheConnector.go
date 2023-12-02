package cacheConnector

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/ssh"
	"github.com/redis/go-redis/v9"
)

var ctx = conf.CTX

func InitEntityCacheConnector() (entityCacheConnector *EntityCacheConnector) {
	entityCacheConnector = new(EntityCacheConnector)
	entityCacheConnector.Init()
	return entityCacheConnector
}

func (self *EntityCacheConnector) Init() *EntityCacheConnector {
	self.setEntitySSH().setEntityConfigRedis().setCacheRedis().pingCacheRedis()
	return self
}

func (self *EntityCacheConnector) setEntitySSH() *EntityCacheConnector {
	isNeedSSH := conf.IsNeedSSH()
	self.IsNeedSSH = isNeedSSH
	if !isNeedSSH {
		return self
	}
	entitySSH := ssh.InitEntitySSH()
	self.EntitySSH = entitySSH
	return self
}

func (self *EntityCacheConnector) setEntityConfigRedis() *EntityCacheConnector {
	ec := conf.GetEntityConfigRedis()
	self.EntityConfigRedis = ec
	return self
}

func (self *EntityCacheConnector) setCacheRedis() *EntityCacheConnector {
	entityRedisOptions := self.initEntityRedisOptions()
	r := redis.NewClient(entityRedisOptions)
	self.CacheRedis = r
	return self
}

func (self *EntityCacheConnector) initEntityRedisOptions() (entityRedisOptions *redis.Options) {
	ec := self.EntityConfigRedis
	redisHost := ec.Host
	redisPort := ec.Port
	redisPassword := ec.Password
	redisDB := ec.DB
	redisNetwork := ssh.NetworkTCP
	maxRetries := conf.MaxRetries
	timeoutDial := conf.TimeoutDial
	timeoutRead := conf.TimeoutRead
	timeoutWrite := conf.TimeoutWrite
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
		timeoutReadDisable := conf.TimeoutReadDisable
		timeoutWriteDisable := conf.TimeoutWriteDisable
		entityRedisOptions.Dialer = funcDialerRedis
		entityRedisOptions.ReadTimeout = timeoutReadDisable
		entityRedisOptions.WriteTimeout = timeoutWriteDisable
	}
	return entityRedisOptions
}

func (self *EntityCacheConnector) pingCacheRedis() *EntityCacheConnector {
	_, errRedisPing := self.CacheRedis.Ping(ctx).Result()
	err.ErrCheck(errRedisPing)
	return self
}

func (self *EntityCacheConnector) CloseCacheConnector() {
	self.closeCacheRedis().closeSSH()
	return
}

func (self *EntityCacheConnector) closeCacheRedis() *EntityCacheConnector {
	if self.CacheRedis == nil {
		return self
	}
	errCloseRedis := self.CacheRedis.Close()
	err.ErrCheck(errCloseRedis)
	self.CacheRedis = nil
	return self
}

func (self *EntityCacheConnector) closeSSH() *EntityCacheConnector {
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

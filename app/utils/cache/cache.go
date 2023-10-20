package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"go.lwh.com/linweihao/lwhFrameGo/app/conf"
	"go.lwh.com/linweihao/lwhFrameGo/app/factory/factoryOfSSH"
	_ "go.lwh.com/linweihao/lwhFrameGo/app/utils/dd"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/err"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/ssh"
	"time"
	// ttt "go.lwh.com/linweihao/lwhFrameGo/app/utils/time"
)

var MaxRetries int = 0
var TimeoutDial time.Duration = time.Second * time.Duration(5)
var TimeoutRead time.Duration = time.Second * time.Duration(1)
var TimeoutWrite time.Duration = time.Second * time.Duration(5)

type EntityCache struct {
	CacheRedis *redis.Client
	EntitySSH  *ssh.EntitySSH
}

////
// Init
////

func InitCache() (entityCache *EntityCache) {
	cacheRedis := &redis.Client{}
	entitySSH := getEntitySSH()
	r := conf.GetEntityConfigRedis()
	cacheRedis = initCacheRedis(r, entitySSH)
	// ttt.ShowTimeAndMsg("PING")
	// resRedisPing, errRedisPing := cacheRedis.Ping().Result()
	_, errRedisPing := cacheRedis.Ping().Result()
	err.ErrCheck(errRedisPing)
	// ttt.ShowTimeAndMsg(resRedisPing)
	entityCache = &EntityCache{
		CacheRedis: cacheRedis,
		EntitySSH:  entitySSH,
	}
	return entityCache
}

func getEntitySSH() (entitySSH *ssh.EntitySSH) {
	entitySSH = nil
	isNeedSSH := isNeedSSH()
	if isNeedSSH {
		// entitySSH = ssh.InitSSHForRedis()
		entitySSH = factoryOfSSH.MakeEntityOfSSHForRedis()
	}
	return entitySSH
}

func isNeedSSH() (isNeedSSH bool) {
	isNeedSSH = conf.IsNeedSSH()
	return isNeedSSH
}

func initCacheRedis(r *conf.EntityConfigRedis, entitySSH *ssh.EntitySSH) (cacheRedis *redis.Client) {
	redisHost := r.Host
	redisPort := r.Port
	redisPassword := r.Password
	redisDB := r.DB
	redisAddr := fmt.Sprintf(
		"%s:%s",
		redisHost,
		redisPort)
	maxRetries := MaxRetries
	timeoutDial := TimeoutDial
	timeoutRead := TimeoutRead
	timeoutWrite := TimeoutWrite
	entityRedisOptions := redis.Options{
		Addr:         redisAddr,
		Password:     redisPassword,
		DB:           redisDB,
		MaxRetries:   maxRetries,
		DialTimeout:  timeoutDial,
		ReadTimeout:  timeoutRead,
		WriteTimeout: timeoutWrite,
	}
	isNeedSSH := isNeedSSH()
	if isNeedSSH {
		entitySSH.SetAddress(redisAddr)
		funcDialerRedis := entitySSH.DialForRedis
		entityRedisOptions.Dialer = funcDialerRedis
	}
	// dd.DD(entityRedisOptions)
	cacheRedis = redis.NewClient(&entityRedisOptions)
	// dd.DD(cacheRedis)
	// ttt.ShowTimeAndMsg("Redis connect success")
	return cacheRedis
}

////
// Close
////

func (self *EntityCache) CloseCache() {
	cacheRedis := self.CacheRedis
	if cacheRedis != nil {
		cacheRedis.Close()
	}
	self.CacheRedis = nil
	// ttt.ShowTimeAndMsg("Cache close success")
	isNeedSSH := conf.IsNeedSSH()
	if isNeedSSH {
		entitySSH := self.EntitySSH
		if entitySSH != nil {
			entitySSH.CloseSSH()
			// ttt.ShowTimeAndMsg("SSH close success")
		}
	}
	self.EntitySSH = nil
	return
}

////
// Exec
////

func (self *EntityCache) SetToCache(cacheKey string, cacheValue string, ttl time.Duration) {
	cacheRedis := self.CacheRedis
	errRedisSet := cacheRedis.Set(cacheKey, cacheValue, ttl).Err()
	err.ErrCheck(errRedisSet)
	return
}

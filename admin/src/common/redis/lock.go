package redis

import (
	"errors"
	"fmt"
	"time"

	"siteol.com/smart/src/common/log"

	"github.com/go-redis/redis"
)

var ErrLocked = errors.New("locked")

// Locker 分布锁接口
type Locker interface {
	Lock(key string, expired int) (v UnLocker, err error)
	LockWithRetry(key string, expired int, retry int, sleep int) (v UnLocker, err error)
}

// UnLocker 分布式解锁接口
type UnLocker interface {
	Unlock() error
}

// 分布锁对象实现对象
type redisLock struct {
	client *redis.Client
	exp    int
}

// 分布式解锁实现对象
type redisUnLock struct {
	client *redis.Client
	k, v   string
}

// NewRedisLock 初始化一个分布式锁对象
func NewRedisLock() Locker {
	return &redisLock{
		client: cluster,
		exp:    86400, // 超长失效1天
	}
}

// 分布式锁移除脚本
var deleteScript = redis.NewScript(`
if redis.call("get", KEYS[1]) == ARGV[1] then
	return redis.call("del", KEYS[1])
else
	return 0
end`)

// Unlock 分布式锁移除
func (un *redisUnLock) Unlock() error {
	err := deleteScript.Run(un.client, []string{un.k}, un.v).Err()
	if err != nil {
		log.ErrorF("redis.locker: unlock %s error: %s", un.k, err)
		return fmt.Errorf("unlock: %s, %s", un.k, err)
	}
	return nil
}

// Lock 获取分布式锁
func (rl *redisLock) Lock(key string, expired int) (v UnLocker, err error) {
	if expired == 0 {
		expired = rl.exp
	}
	ret := rl.client.Do("SET", key, "LOCK", "ex", expired, "nx")
	err = ret.Err()
	if err == redis.Nil {
		err = ErrLocked
		return
	}
	if err != nil {
		return nil, err
	}
	v = &redisUnLock{k: key, v: "LOCK", client: rl.client}
	return
}

// LockWithRetry 多次尝试的分布式锁获取
func (rl *redisLock) LockWithRetry(key string, expired int, retry int, sleep int) (v UnLocker, err error) {
	for i := 0; i < retry; i++ {
		v, err = rl.Lock(key, expired)
		if err == ErrLocked {
			sleepDuration := time.Duration(sleep)
			time.Sleep(sleepDuration * time.Millisecond) // 绑定毫秒
			continue
		}
		return
	}
	err = ErrLocked
	return
}

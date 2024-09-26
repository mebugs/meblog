package redis

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"siteol.com/smart/src/common/log"

	"github.com/go-redis/redis"
)

var ErrNotFound = errors.New("not found")
var DuplicateKeys = errors.New("Duplicate Keys")

// Set 基础的缓存设置 超时为0表示永不超时
func Set(key string, obj any, millisecond uint64) (err error) {
	return SetByTimeDuration(key, obj, time.Millisecond*time.Duration(millisecond))
}

// SetTTL 仅更新换群时间
func SetTTL(key string, millisecond uint64) (err error) {
	cmd := cluster.PExpire(key, time.Millisecond*time.Duration(millisecond))
	return cmd.Err()
}

// SetByTimeDuration 基础的缓存设置 根据传入的时间设置值
func SetByTimeDuration(key string, obj any, timeDuration time.Duration) (err error) {
	var value string
	switch v := obj.(type) {
	case string:
		value = v
	case []byte:
		value = string(v)
	case int, int8, int16, int32, int64:
		value = fmt.Sprintf("%d", v)

	default:
		bs, e := json.Marshal(obj)
		if e != nil {
			err = e
			return
		}
		value = string(bs)
	}
	err = cluster.Set(key, value, timeDuration).Err()
	return
}

// SetNX 分布式设值 超时为0表示永不超时
func SetNX(key string, obj any, millisecond uint64) (err error) {
	var value string
	switch v := obj.(type) {
	case string:
		value = v
	case []byte:
		value = string(v)
	case int, int8, int16, int32, int64:
		value = fmt.Sprintf("%d", v)

	default:
		bs, e := json.Marshal(obj)
		if e != nil {
			err = e
			return
		}
		value = string(bs)
	}
	res, err := cluster.SetNX(key, value, time.Millisecond*time.Duration(millisecond)).Result()
	if err != nil {
		log.Error(err)
		return err
	} else if !res {
		return DuplicateKeys
	}
	return
}

// Get 获取缓存
func Get(key string) (ret string, err error) {
	v := cluster.Get(key)
	if v.Err() == redis.Nil {
		err = ErrNotFound
		return
	}
	if v.Err() != nil {
		err = v.Err()
		return
	}
	ret = v.Val()
	return
}

// GetTTL 获取缓存超期时间
func GetTTL(key string) (ret time.Duration, err error) {
	v := cluster.TTL(key)
	if v.Err() == redis.Nil {
		err = ErrNotFound
		return
	}
	if v.Err() != nil {
		err = v.Err()
		return
	}
	ret = v.Val()
	return
}

// Del 移除缓存
func Del(key string) (err error) {
	err = cluster.Del(key).Err()
	return
}

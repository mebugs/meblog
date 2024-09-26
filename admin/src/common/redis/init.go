package redis

import (
	"os"

	"siteol.com/smart/src/common/log"
	"siteol.com/smart/src/config"

	"github.com/go-redis/redis"
)

var cluster *redis.Client

// Init Redis初始化
func Init(traceId string) {
	// redis
	cluster = redis.NewClient(&redis.Options{
		Addr:     config.JsonConfig.Redis.Addr,
		DB:       config.JsonConfig.Redis.DB,
		Password: config.JsonConfig.Redis.Password,
	})
	err := cluster.Ping().Err()
	if err != nil {
		log.ErrorTF(traceId, "Init Redis Fail . Err : %s", err)
		os.Exit(1)
	} else {
		log.InfoTF(traceId, "Init Redis success")
	}
}

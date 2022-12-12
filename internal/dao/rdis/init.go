package rdis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/mangohow/cloud-ide-webserver/conf"
	"time"
)

var client *redis.Client

func InitRedis() error {
	client = redis.NewClient(&redis.Options{
		Addr:         conf.RedisConfig.Addr,
		Password:     conf.RedisConfig.Password,
		DB:           int(conf.RedisConfig.DB),
		PoolSize:     int(conf.RedisConfig.PoolSize),     // 连接池最大socket连接数
		MinIdleConns: int(conf.RedisConfig.MinIdleConns), // 最少连接维持数
	})

	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, err := client.Ping(timeoutCtx).Result()
	if err != nil {
		return err
	}

	return nil
}

func CloseRedisConn() {
	client.Close()
}

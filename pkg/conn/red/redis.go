package red

import (
	"context"
	//"log"
	"github.com/redis/go-redis/v9"
	"preject/pkg/log"
)

func ConnRedis(addr, password string, db int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,     // Redis服务器地址
		Password: password, // 密码（如果有的话）
		DB:       db,       // 使用默认的数据库
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Logger.Error("redis_db connect ping failed, err:", err)
		return client, err
	}
	log.Logger.Info("连接成功:", pong)
	return client, nil
}

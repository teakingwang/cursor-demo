package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/teakingwang/cursor-demo/config"
)

var (
	Client *redis.Client
	Ctx    = context.Background()
)

func InitRedis() error {
	redisConfig := config.GlobalConfig.Redis
	Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
		PoolSize: redisConfig.PoolSize,
	})

	// 测试连接
	if err := Client.Ping(Ctx).Err(); err != nil {
		return fmt.Errorf("Redis连接失败: %v", err)
	}

	return nil
}

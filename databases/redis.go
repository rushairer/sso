package databases

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisClient 定义了 Redis 操作的接口
type RedisClient interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Get(ctx context.Context, key string) *redis.StringCmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
}

// redisClientImpl 是 RedisClient 接口的实现
type redisClientImpl struct {
	client *redis.Client
}

// NewRedisClient 创建一个新的 Redis 客户端
func NewRedisClient(client *redis.Client) RedisClient {
	return &redisClientImpl{client: client}
}

// Set 实现 RedisClient 接口的 Set 方法
func (r *redisClientImpl) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return r.client.Set(ctx, key, value, expiration)
}

// Get 实现 RedisClient 接口的 Get 方法
func (r *redisClientImpl) Get(ctx context.Context, key string) *redis.StringCmd {
	return r.client.Get(ctx, key)
}

// Del 实现 RedisClient 接口的 Del 方法
func (r *redisClientImpl) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	return r.client.Del(ctx, keys...)
}

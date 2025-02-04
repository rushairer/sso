package databases

import (
	"context"
	"database/sql"
	"log"
	"sync"

	"github.com/redis/go-redis/v9"
	"github.com/rushairer/sso/config"
	"github.com/rushairer/sso/utils/errors"
)

var (
	db          *sql.DB
	redisClient *redis.Client
	once        sync.Once
)

// InitDB 初始化数据库连接
func InitDB() (*sql.DB, error) {
	var err error
	once.Do(func() {
		db, err = sql.Open("mysql", config.MySQLDSN)
		if err != nil {
			log.Printf("Failed to connect to database: %v\n", err)
			return
		}

		err = db.Ping()
		if err != nil {
			log.Printf("Failed to ping database: %v\n", err)
			return
		}
	})

	if err != nil {
		return nil, errors.NewInternalError("Database connection failed", err)
	}

	return db, nil
}

// InitRedis 初始化Redis连接
func InitRedis() (RedisClient, error) {
	var err error
	once.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr: config.RedisDSN,
			DB:   0,
		})

		// 测试连接
		_, err = redisClient.Ping(context.Background()).Result()
		if err != nil {
			log.Printf("Failed to connect to Redis: %v\n", err)
			return
		}
	})

	if err != nil {
		return nil, errors.NewInternalError("Redis connection failed", err)
	}

	return NewRedisClient(redisClient), nil
}

// CloseDB 关闭数据库连接
func CloseDB() {
	if db != nil {
		db.Close()
	}
}

// CloseRedis 关闭Redis连接
func CloseRedis() {
	if redisClient != nil {
		redisClient.Close()
	}
}

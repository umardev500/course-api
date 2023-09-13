package storage

import (
	"context"
	"course-api/config"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

type RedisStorage struct {
	storage *redis.Client
}

func NewRedisStorage() fiber.Storage {
	return &RedisStorage{
		storage: config.NewRedis(),
	}
}

func (rs *RedisStorage) Get(key string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return rs.storage.Get(ctx, key).Bytes()
}

func (rs *RedisStorage) Set(key string, val []byte, exp time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return rs.storage.Set(ctx, key, val, exp).Err()
}

func (rs *RedisStorage) Delete(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return rs.storage.Del(ctx, key).Err()
}

func (rs *RedisStorage) Reset() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return rs.storage.FlushAll(ctx).Err()
}

func (rs *RedisStorage) Close() error {
	return rs.storage.Close()
}

package config

import (
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func NewRedis() *redis.Client {
	dbIndex, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	rdb := redis.NewClient(&redis.Options{
		Addr:     "course_redis:6379",
		Password: "",
		DB:       dbIndex,
	})

	return rdb
}

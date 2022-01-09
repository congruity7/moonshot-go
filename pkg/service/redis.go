package service

import (
	"github.com/go-redis/redis"
)

type RedisService struct {
	rc *redis.Client
}

func NewRedisService(rc *redis.Client) *RedisService {
	return &RedisService{rc: rc}
}

package service

import (
	"github.com/go-redis/redis"
)

type RedisService struct {
	Client *redis.Client
}

func NewRedisService(client *redis.Client) *RedisService {
	return &RedisService{Client: client}
}

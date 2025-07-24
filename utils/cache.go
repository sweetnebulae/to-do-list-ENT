package utils

import (
	"context"
	"github.com/redis/go-redis/v9"
	"os"
	"time"
)

type CacheService struct {
	client *redis.Client
}

var ctx = context.Background()

func NewCacheService() *CacheService {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "",
		DB:       0,
	})
	return &CacheService{client: client}
}

func (s *CacheService) Set(key string, value string, expiration time.Duration) error {
	return s.client.Set(ctx, key, value, expiration).Err()
}

func (s *CacheService) Get(key string) (string, error) {
	return s.client.Get(ctx, key).Result()
}

func (s *CacheService) Delete(key string) error {
	return s.client.Del(ctx, key).Err()
}

func (s *CacheService) StoreSession(key string, value string, expiration time.Duration) error {
	return s.client.Set(ctx, key, value, expiration).Err()
}

func (s *CacheService) GetSession(key string) (string, error) {
	return s.client.Get(ctx, key).Result()
}

func (s *CacheService) DeleteSession(key string) error {
	return s.client.Del(ctx, key).Err()
}

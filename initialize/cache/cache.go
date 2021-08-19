package cache

import (
	"context"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/snowlyg/go-tenancy/g"
)

func Cache() redis.UniversalClient {
	universalOptions := &redis.UniversalOptions{
		Addrs:       strings.Split(g.TENANCY_CONFIG.Redis.Addr, ","),
		Password:    g.TENANCY_CONFIG.Redis.Password,
		PoolSize:    g.TENANCY_CONFIG.Redis.PoolSize,
		IdleTimeout: 300 * time.Second,
	}
	return redis.NewUniversalClient(universalOptions)
}

func SetCache(key string, value interface{}, expiration time.Duration) error {
	err := g.TENANCY_CACHE.Set(context.Background(), key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetCacheString(key string) (string, error) {
	return g.TENANCY_CACHE.Get(context.Background(), key).Result()
}

func DeleteCache(key string) (int64, error) {
	return g.TENANCY_CACHE.Del(context.Background(), key).Result()
}

func GetCacheBytes(key string) ([]byte, error) {
	return g.TENANCY_CACHE.Get(context.Background(), key).Bytes()
}
func GetCacheUint(key string) (uint64, error) {
	return g.TENANCY_CACHE.Get(context.Background(), key).Uint64()
}

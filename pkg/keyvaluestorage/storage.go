package keyvaluestorage

import (
	"context"
	"time"

	redis "github.com/go-redis/redis/v8"
)

type Storage struct {
	rdb *redis.Client
}

var ctx = context.Background()

func New(redisURL string) *Storage {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	storage := &Storage{rdb}

	return storage
}

func (s *Storage) Get(key string) (string, error) {
	return s.rdb.Get(ctx, key).Result()
}

func (s *Storage) Set(key, value string, expire *time.Duration) error {
	var expiration time.Duration

	if expire != nil {
		expiration = *expire
	}

	return s.rdb.Set(ctx, key, value, expiration).Err()
}

func (s *Storage) Remove(key string) error {
	return s.rdb.Del(ctx, key).Err()
}

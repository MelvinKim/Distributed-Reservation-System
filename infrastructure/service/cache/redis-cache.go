package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/MelvinKim/Hotel-Reservation-System/domain"
	"github.com/go-redis/redis/v8"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, exp time.Duration) *redisCache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}

func (cache *redisCache) Set(key string, value *domain.Guest) {
	ctx := context.Background()
	// get client, use it to assign the key value pair
	client := cache.getClient()

	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	client.Set(ctx, key, string(json), cache.expires*time.Second)
}

func (cache *redisCache) Get(key string) *domain.Guest {
	ctx := context.Background()
	client := cache.getClient()

	val, err := client.Get(ctx, key).Result()
	if err != nil {
		return nil
	}

	guest := domain.Guest{}
	err = json.Unmarshal([]byte(val), &guest)
	if err != nil {
		return nil
	}
	return &guest
}

package redisadapter

import (
	"context"
	"sync"

	"github.com/muhfaris/rocket-hexagonal/internal/core/port/outbound/repository"
	"github.com/redis/go-redis/v9"
)

type RedisOptions struct {
	Addr     string
	Username string
	Password string
	DB       int
}

var (
	cache *redis.Client
	once  sync.Once
)

func New(opt RedisOptions) repository.CacheRepository {
	if cache != nil {
		return nil
	}

	once.Do(func() {
		rdb := redis.NewClient(&redis.Options{
			Addr:     opt.Addr,
			Password: opt.Password, // no password set
			Username: opt.Username,
			DB:       opt.DB, // use default DB
		})

		if err := rdb.Ping(context.Background()).Err(); err != nil {
			panic(err)
		}

		cache = rdb
	})

	return &Client{
		conn: cache,
	}
}

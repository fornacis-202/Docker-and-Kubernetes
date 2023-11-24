package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// Storage manages the database connections.
type Redis struct {
	timeout int
	rdb     *redis.Client
}

type Config struct {
	Host    string `koanf:"host"`
	Timeout int    `koanf:"timeout"`
}

// New builds a new storage.
func New(cfg Config) *Redis {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Host,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := rdb.Ping(context.Background()).Result()
	fmt.Println(pong, err)

	return &Redis{
		timeout: cfg.Timeout,
		rdb:     rdb,
	}

}

// Read key from database.
func (s *Redis) Read(key string) (string, error) {
	// create context
	ctx := context.Background()
	fmt.Println(s.rdb.Get(ctx, key).Result())
	return s.rdb.Get(ctx, key).Result()
}

// Write a set into database.
func (s *Redis) Write(key string, value string) error {
	// create context
	ctx := context.Background()
	return s.rdb.Set(ctx, key, value, time.Duration(s.timeout)*time.Minute).Err()
}

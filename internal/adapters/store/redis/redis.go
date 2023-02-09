package redis_store

import (
	"context"
	"ddd-template/internal/adapters/store"
	"ddd-template/internal/config"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

// NewRedisStore 创建基于redis存储实例
func NewRedisStore(cfg *config.Configs) store.Store {
	cli := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		DB:       cfg.Redis.DB,
		Password: cfg.Redis.Password,
	})
	if cfg.Redis.Use {
		if err := cli.Ping(context.Background()).Err(); err != nil {
			log.Fatalln(err)
		}
	}
	return &Store{
		cli:    cli,
		prefix: cfg.Redis.KeyPrefix,
	}
}

// Store redis存储
type Store struct {
	cli    *redis.Client
	prefix string
}

func (s *Store) wrapperKey(key string) string {
	return fmt.Sprintf("%s|%s", s.prefix, key)
}

func (s *Store) Get(ctx context.Context, key string) (string, error) {
	return s.cli.Get(ctx, s.wrapperKey(key)).Result()
}

// Set ...
func (s *Store) Set(ctx context.Context, uuid string, val string, expiration time.Duration) error {
	cmd := s.cli.Set(ctx, s.wrapperKey(uuid), val, expiration)
	return cmd.Err()
}

// Delete ...
func (s *Store) Delete(ctx context.Context, tokenString string) (bool, error) {
	cmd := s.cli.Del(ctx, s.wrapperKey(tokenString))
	if err := cmd.Err(); err != nil {
		return false, err
	}
	return cmd.Val() > 0, nil
}

// Check ...
func (s *Store) Check(ctx context.Context, tokenString string) (bool, error) {
	cmd := s.cli.Exists(ctx, s.wrapperKey(tokenString))
	if err := cmd.Err(); err != nil {
		return false, err
	}
	return cmd.Val() > 0, nil
}

func (s *Store) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return s.cli.SetNX(ctx, s.wrapperKey(key), value, expiration).Err()
}

func (s *Store) Del(ctx context.Context, key string) error {
	return s.cli.Del(ctx, s.wrapperKey(key)).Err()
}

func (s *Store) DelByKeyPrefix(ctx context.Context, keyPrefix string) error {
	var cursor uint64
	var keys []string
	var err error

	for {
		keys, cursor, err = s.cli.Scan(ctx, cursor, keyPrefix+"*", 100).Result()
		if err != nil {
			return err
		}

		if len(keys) > 0 {
			if err = s.cli.Del(ctx, keys...).Err(); err != nil {
				return err
			}
		}

		if cursor == 0 {
			break
		}
	}

	return nil
}

func (s *Store) Exists(ctx context.Context, key string) bool {
	return s.cli.Exists(ctx, s.wrapperKey(key)).Val() == 1
}

// Close ...
func (s *Store) Close() error {
	return s.cli.Close()
}

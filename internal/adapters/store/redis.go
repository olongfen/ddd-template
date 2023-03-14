package store

import (
	"context"
	"ddd-template/internal/rely"
	"fmt"
	"go.uber.org/zap"
	"time"

	"github.com/go-redis/redis/v8"
)

// NewRedisStore 创建基于redis存储实例
func NewRedisStore(cfg *rely.Configs, logger *zap.Logger) (ret Store, fc func(), err error) {
	cli := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		DB:       cfg.Redis.DB,
		Password: cfg.Redis.Password,
	})
	if cfg.Redis.Use {
		if err = cli.Ping(context.Background()).Err(); err != nil {
			return
		}
	}
	sto := &storeRedis{
		cli:    cli,
		prefix: cfg.Redis.KeyPrefix,
	}
	return sto, func() {
		logger.Info("redis client close")
		_ = sto.Close()
	}, nil
}

// storeRedis redis存储
type storeRedis struct {
	cli    *redis.Client
	prefix string
}

func (s *storeRedis) wrapperKey(key string) string {
	return fmt.Sprintf("%s|%s", s.prefix, key)
}

func (s *storeRedis) Get(ctx context.Context, key string) (string, error) {
	return s.cli.Get(ctx, s.wrapperKey(key)).Result()
}

// Set ...
func (s *storeRedis) Set(ctx context.Context, uuid string, val string, expiration time.Duration) error {
	cmd := s.cli.Set(ctx, s.wrapperKey(uuid), val, expiration)
	return cmd.Err()
}

// Delete ...
func (s *storeRedis) Delete(ctx context.Context, tokenString string) (bool, error) {
	cmd := s.cli.Del(ctx, s.wrapperKey(tokenString))
	if err := cmd.Err(); err != nil {
		return false, err
	}
	return cmd.Val() > 0, nil
}

// Check ...
func (s *storeRedis) Check(ctx context.Context, tokenString string) (bool, error) {
	cmd := s.cli.Exists(ctx, s.wrapperKey(tokenString))
	if err := cmd.Err(); err != nil {
		return false, err
	}
	return cmd.Val() > 0, nil
}

func (s *storeRedis) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return s.cli.SetNX(ctx, s.wrapperKey(key), value, expiration).Err()
}

func (s *storeRedis) Del(ctx context.Context, key string) error {
	return s.cli.Del(ctx, s.wrapperKey(key)).Err()
}

func (s *storeRedis) DelByKeyPrefix(ctx context.Context, keyPrefix string) error {
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

func (s *storeRedis) Exists(ctx context.Context, key string) bool {
	return s.cli.Exists(ctx, s.wrapperKey(key)).Val() == 1
}

// Close ...
func (s *storeRedis) Close() error {
	return s.cli.Close()
}

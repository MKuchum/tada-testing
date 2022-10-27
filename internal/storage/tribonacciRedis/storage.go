package tribonacciRedis

import (
	"context"
	"encoding/json"
	"github.com/MKuchum/tada-testing/internal/storage"
	"github.com/go-redis/redis/v9"
	"go.uber.org/zap"
)

type Storage struct {
	r      *redis.Client
	logger *zap.Logger
}

func NewTribonacciStorageRedis(addr string, password string, db int, logger *zap.Logger) (storage.TribonacciStorage, error) {
	r := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	if err := r.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	logger.Info("Successfully ping redis")
	return &Storage{r: r, logger: logger}, nil
}

func (s *Storage) genKey(signature []float32) (string, error) {
	if res, err := json.Marshal(signature); err != nil {
		return "", err
	} else {
		return string(res), nil
	}
}

func (s *Storage) get(signature []float32) ([]float32, error) {
	key, err := s.genKey(signature)
	if err != nil {
		return nil, err
	}
	val, err := s.r.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	} else {
		var res []float32
		if err := json.Unmarshal([]byte(val), &res); err != nil {
			return nil, err
		}
		return res, nil
	}
}

func (s *Storage) set(signature []float32, values []float32) error {
	key, err := s.genKey(signature)
	if err != nil {
		return err
	}
	value, err := json.Marshal(values)
	if err != nil {
		return err
	}
	return s.r.Set(context.Background(), key, string(value), 0).Err()
}

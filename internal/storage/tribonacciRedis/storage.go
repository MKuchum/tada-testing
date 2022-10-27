package tribonacciRedis

import (
	"context"
	"encoding/json"
	"github.com/MKuchum/tada-testing/internal/storage"
	"github.com/go-redis/redis/v9"
)

type Storage struct {
	r *redis.Client
}

func NewTribonacciStorageRedis(addr string, password string, db int) storage.TribonacciStorage {
	return &Storage{
		r: redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       db,
		}),
	}
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

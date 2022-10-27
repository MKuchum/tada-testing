package tribonacci

import (
	"github.com/MKuchum/tada-testing/internal/storage"
	"github.com/MKuchum/tada-testing/internal/storage/tribonacciInMemory"
	"github.com/MKuchum/tada-testing/internal/storage/tribonacciRedis"
	"go.uber.org/zap"
)

type Tribonacci struct {
	s      storage.TribonacciStorage
	logger *zap.Logger
}

func NewTribonacciRedis(addr string, password string, db int, logger *zap.Logger) (*Tribonacci, error) {
	s, err := tribonacciRedis.NewTribonacciStorageRedis(addr, password, db, logger)
	if err != nil {
		return nil, err
	}
	return &Tribonacci{s: s, logger: logger}, nil
}

func NewTribonacciInMemory() *Tribonacci {
	return &Tribonacci{s: tribonacciInMemory.NewTribonacciStorageInMemory()}
}

package tribonacci

import (
	"github.com/MKuchum/tada-testing/internal/storage"
	"github.com/MKuchum/tada-testing/internal/storage/tribonacciInMemory"
	"github.com/MKuchum/tada-testing/internal/storage/tribonacciRedis"
)

type Tribonacci struct {
	s storage.TribonacciStorage
}

func NewTribonacciRedis(addr string, password string, db int) (*Tribonacci, error) {
	s, err := tribonacciRedis.NewTribonacciStorageRedis(addr, password, db)
	if err != nil {
		return nil, err
	}
	return &Tribonacci{s: s}, nil
}

func NewTribonacciInMemory() *Tribonacci {
	return &Tribonacci{s: tribonacciInMemory.NewTribonacciStorageInMemory()}
}

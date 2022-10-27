package tribonacci

import (
	"github.com/MKuchum/tada-testing/internal/storage"
	"github.com/MKuchum/tada-testing/internal/storage/tribonacciInMemory"
	"github.com/MKuchum/tada-testing/internal/storage/tribonacciRedis"
)

type Tribonacci struct {
	s storage.TribonacciStorage
}

func NewTribonacciRedis(addr string, password string, db int) *Tribonacci {
	return &Tribonacci{s: tribonacciRedis.NewTribonacciStorageRedis(addr, password, db)}
}

func NewTribonacciInMemory() *Tribonacci {
	return &Tribonacci{s: tribonacciInMemory.NewTribonacciStorageInMemory()}
}

package tribonacci

import (
	"github.com/MKuchum/tada-testing/internal/storage"
	"github.com/MKuchum/tada-testing/internal/storage/tribonacciInMemory"
)

type Tribonacci struct {
	s storage.TribonacciStorage
}

func NewTribonacci() *Tribonacci {
	return &Tribonacci{s: tribonacciInMemory.NewTribonacciStorageInMemory()}
}

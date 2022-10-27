package tribonacciInMemory

import (
	"fmt"
	"github.com/MKuchum/tada-testing/internal/storage"
	"strings"
)

type Storage struct {
	m map[string][]float32
}

func NewTribonacciStorageInMemory() storage.TribonacciStorage {
	return &Storage{
		m: make(map[string][]float32),
	}
}

func (s *Storage) genKey(signature []float32) string {
	strs := make([]string, 0, len(signature))
	for _, v := range signature {
		strs = append(strs, fmt.Sprintf("%f", v))
	}
	return strings.Join(strs, "#")
}

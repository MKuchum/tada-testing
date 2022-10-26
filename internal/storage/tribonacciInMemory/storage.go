package tribonacciInMemory

import "github.com/MKuchum/tada-testing/internal/storage"

type Storage struct {
}

func NewTribonacciStorageInMemory() storage.TribonacciStorage {
	return &Storage{}
}

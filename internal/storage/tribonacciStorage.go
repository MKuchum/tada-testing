package storage

type TribonacciStorage interface {
	Get(signature []float32, n int64) ([]float32, int, error) // slice, len, err
	Set(signature []float32, n int64, values []float32) error
}

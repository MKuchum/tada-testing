package storage

type TribonacciStorage interface {
	Get(signature []float32, n int64) ([]float32, error) // slice, err
	Set(signature []float32, values []float32) error
}

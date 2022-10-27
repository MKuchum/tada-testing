package storage

type TribonacciStorage interface {
	Get(signature []float32, n int) ([]float32, error)
	Set(signature []float32, values []float32) error
}

package tribonacciInMemory

import "github.com/MKuchum/tada-testing/models"

func (s *Storage) Get(signature []float32, n int64) ([]float32, error) {
	if len(signature) != 3 {
		return nil, models.WrongSignatureLenErr
	}
	if n <= 0 {
		return nil, models.WrongSequenceLenErr
	}

	key := s.genKey(signature)
	if v, ok := s.m[key]; ok {
		if len(v) >= int(n) {
			return v[:n], nil
		} else {
			return v, nil
		}
	} else {
		if n <= 3 {
			return signature[:n], nil
		} else {
			return signature, nil
		}
	}
}

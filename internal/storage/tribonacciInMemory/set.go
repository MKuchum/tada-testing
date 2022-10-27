package tribonacciInMemory

import "github.com/MKuchum/tada-testing/models"

func (s *Storage) Set(signature []float32, values []float32) error {
	if len(signature) != 3 {
		return models.WrongSignatureLenErr
	}

	key := s.genKey(signature)
	if v, ok := s.m[key]; ok {
		if len(values) > len(v) {
			s.m[key] = values
		}
	} else {
		s.m[key] = values
	}
	return nil
}

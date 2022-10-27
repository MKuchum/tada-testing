package tribonacciRedis

import "github.com/MKuchum/tada-testing/models"

func (s *Storage) Set(signature []float32, values []float32) error {
	if len(signature) != 3 {
		return models.WrongSignatureLenErr
	}

	v, err := s.get(signature)
	if err != nil {
		return err
	}
	if v != nil {
		if len(values) > len(v) {
			return s.set(signature, values)
		}
	} else {
		return s.set(signature, values)
	}
	return nil
}

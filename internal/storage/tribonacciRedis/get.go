package tribonacciRedis

import (
	"github.com/MKuchum/tada-testing/models"
	"go.uber.org/zap"
)

func (s *Storage) Get(signature []float32, n int) ([]float32, error) {
	if len(signature) != 3 {
		return nil, models.WrongSignatureLenErr
	}
	if n <= 0 {
		return nil, models.WrongSequenceLenErr
	}

	v, err := s.get(signature)
	if err != nil {
		return nil, err
	}
	s.logger.Info("redis get", zap.Any("signature", signature), zap.Any("response", v))
	if v != nil {
		if len(v) >= n {
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

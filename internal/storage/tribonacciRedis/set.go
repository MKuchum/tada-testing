package tribonacciRedis

import (
	"github.com/MKuchum/tada-testing/models"
	"go.uber.org/zap"
)

func (s *Storage) Set(signature []float32, values []float32) error {
	if len(signature) != 3 {
		return models.WrongSignatureLenErr
	}

	if len(values) < len(signature) {
		values = signature
	}

	v, err := s.get(signature)
	if err != nil {
		return err
	}
	if v != nil {
		if len(values) > len(v) {
			if err := s.set(signature, values); err != nil {
				return err
			}
			s.logger.Info("redis set", zap.Any("signature", signature), zap.Any("values", values))
			return nil
		}
	} else {
		if err := s.set(signature, values); err != nil {
			return err
		}
		s.logger.Info("redis set", zap.Any("signature", signature), zap.Any("values", values))
		return nil
	}
	return nil
}

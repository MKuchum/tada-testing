package tribonacci

import (
	"github.com/MKuchum/tada-testing/models"
	"go.uber.org/zap"
)

func (t *Tribonacci) Generate(input *models.TribonacciInput) (*models.TribonacciOutput, error) {
	t.logger.Info("start generate", zap.Any("input", input))
	values, err := t.s.Get(input.Signature, input.N)
	if err != nil {
		return nil, err
	}
	for len(values) < input.N {
		newValue := values[len(values)-1] + values[len(values)-2] + values[len(values)-3]
		values = append(values, newValue)
	}
	if err := t.s.Set(input.Signature, values); err != nil {
		return nil, err
	}
	output := &models.TribonacciOutput{Values: values}
	t.logger.Info("end generate", zap.Any("input", input), zap.Any("output", output))
	return output, nil
}

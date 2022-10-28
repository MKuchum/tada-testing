package tribonacci

import (
	"github.com/MKuchum/tada-testing/models"
	"go.uber.org/zap"
)

func (t *Tribonacci) Generate(input *models.TribonacciInput) (*models.TribonacciOutput, error) {
	t.logger.Info("start generate", zap.Any("input", input))
	sequence, err := t.s.Get(input.Signature, input.N)
	if err != nil {
		return nil, err
	}
	for len(sequence) < input.N {
		newValue := sequence[len(sequence)-1] + sequence[len(sequence)-2] + sequence[len(sequence)-3]
		sequence = append(sequence, newValue)
	}
	if err := t.s.Set(input.Signature, sequence); err != nil {
		return nil, err
	}
	output := &models.TribonacciOutput{Sequence: sequence}
	t.logger.Info("end generate", zap.Any("input", input), zap.Any("output", output))
	return output, nil
}

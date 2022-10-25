package tribonacci

import "github.com/MKuchum/tada-testing/models"

func Tribonacci(input *models.TribonacciInput) (*models.TribonacciOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	values := make([]float32, 3, input.N)
	copy(values, input.Signature)
	for int64(len(values)) < input.N {
		newValue := values[len(values)-1] + values[len(values)-2] + values[len(values)-3]
		values = append(values, newValue)
	}
	return &models.TribonacciOutput{Values: values[:input.N]}, nil
}

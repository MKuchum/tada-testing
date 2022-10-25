package Tribonacci

import "github.com/MKuchum/tada-testing/models"

func Tribonacci(input *models.TribonacciInput) (*models.TribonacciOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, err
	}
	values := make([]float32, 3, input.GetN())
	copy(values, input.GetSignature())
	for int64(len(values)) < input.GetN() {
		newValue := values[len(values)-1] + values[len(values)-2] + values[len(values)-3]
		values = append(values, newValue)
	}
	return models.NewTribonacciOutput(values[:input.GetN()]), nil
}

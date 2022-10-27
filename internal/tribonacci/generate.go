package tribonacci

import (
	"github.com/MKuchum/tada-testing/models"
	"log"
)

func (t *Tribonacci) Generate(input *models.TribonacciInput) (*models.TribonacciOutput, error) {
	values, err := t.s.Get(input.Signature, input.N)
	if err != nil {
		return nil, err
	}
	log.Println(values)
	for len(values) < input.N {
		newValue := values[len(values)-1] + values[len(values)-2] + values[len(values)-3]
		values = append(values, newValue)
	}
	if err := t.s.Set(input.Signature, values); err != nil {
		return nil, err
	}
	return &models.TribonacciOutput{Values: values}, nil
}

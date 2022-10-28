package tests

import (
	"github.com/MKuchum/tada-testing/models"
	"strconv"
	"testing"
)

var ts = []*Test{
	{Signature: []float32{0, 0, 0}, N: 10, Sequence: []float32{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
	{Signature: []float32{0, 0, 1}, N: 1, Sequence: []float32{0}},
	{Signature: []float32{0, 0, 1}, N: 2, Sequence: []float32{0, 0}},
	{Signature: []float32{0, 0, 1}, N: 3, Sequence: []float32{0, 0, 1}},
	{Signature: []float32{0, 0, 1}, N: 10, Sequence: []float32{0, 0, 1, 1, 2, 4, 7, 13, 24, 44}},
	{N: 10, Sequence: []float32{1, 1, 1, 3, 5, 9, 17, 31, 57, 105}},
	{Signature: []float32{-1, -1, -1}, N: 10, Sequence: []float32{-1, -1, -1, -3, -5, -9, -17, -31, -57, -105}},
}

func TestSuccessfully(t *testing.T) {
	for i, test := range ts {
		input := &models.TribonacciInput{Signature: test.Signature, N: test.N}
		output, _, err := doReq(t, input)
		assertError(t, err)
		assertEqualsTribonacci(t, test.Sequence, output.Sequence, "test "+strconv.Itoa(i))
	}
}

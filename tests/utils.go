package tests

import "testing"

func assertError(t *testing.T, err error) {
	if err == nil {
		return
	}
	t.Error(err)
}

func assertEqualsTribonacci(t *testing.T, expected []float32, real []float32, comment ...string) {
	c := ""
	if len(comment) > 0 {
		c = comment[0]
	}
	if len(expected) != len(real) {
		t.Errorf("%s, expected = %v, real = %v", c, expected, real)
	}
	for i := range expected {
		if expected[i] != real[i] {
			t.Errorf("%s, expected = %v, real = %v", c, expected, real)
		}
	}
}

type Test struct {
	Signature []float32
	N         int
	Sequence  []float32
}

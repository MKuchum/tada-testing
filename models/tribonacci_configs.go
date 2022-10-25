package models

type TribonacciInput struct {
	Signature []float32 `json:"signature"`
	N         int64     `json:"n"`
}

func (input *TribonacciInput) Validate() error {
	if len(input.Signature) != 3 {
		return WrongSignatureLenErr
	}
	if input.N <= 0 {
		return WrongSequenceLenErr
	}
	return nil
}

type TribonacciOutput struct {
	Values []float32 `json:"values"`
}

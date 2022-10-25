package models

type TribonacciInput struct {
	signature []float32
	n         int64
}

func NewTribonacciInput(signature []float32, n int64) *TribonacciInput {
	return &TribonacciInput{signature: signature, n: n}
}

func (input *TribonacciInput) Validate() error {
	if len(input.signature) != 3 {
		return WrongSignatureLenErr
	}
	if input.n <= 0 {
		return WrongSequenceLenErr
	}
	return nil
}

func (input *TribonacciInput) GetN() int64 {
	return input.n
}

func (input *TribonacciInput) GetSignature() []float32 {
	return input.signature
}

type TribonacciOutput struct {
	values []float32
}

func NewTribonacciOutput(values []float32) *TribonacciOutput {
	return &TribonacciOutput{values: values}
}

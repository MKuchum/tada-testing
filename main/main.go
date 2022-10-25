package main

import (
	"fmt"
	"github.com/MKuchum/tada-testing/internal/Tribonacci"
	"github.com/MKuchum/tada-testing/models"
)

func main() {
	fmt.Println(Tribonacci.Tribonacci(models.NewTribonacciInput([]float32{0, 0, 1}, 10)))
}

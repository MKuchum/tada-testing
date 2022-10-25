package models

import "fmt"

var WrongSignatureLenErr = fmt.Errorf("the length of the signature must be 3")
var WrongSequenceLenErr = fmt.Errorf("the length of the secuence (n) must be more than 0")

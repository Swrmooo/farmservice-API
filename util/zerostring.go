package util

import (
	lib "github.com/ttoonn112/ktgolib"
)

func ZeroString(value int64, numberOfZero int) string {
	padStr := ""
	for k := 0; k < numberOfZero; k++ {
		padStr += "0"
	}
	thestr := padStr + lib.I64_S(value)
	return thestr[len(thestr)-numberOfZero:]
}

package utils

import (
	"fmt"
	"math/big"
)

func ReverseBigInt(number *big.Int) *big.Int {
	str := number.String()
	revStr := ReverseString(str)
	n, success := big.NewInt(0).SetString(revStr, 10)
	if !success {
		panic(fmt.Sprintf("cannot big.NewInt.SetString(%q)", revStr))
	}
	return n
}

func IsPalindromeBigInt(number *big.Int) bool {
	rev := ReverseBigInt(number)
	return number.String() == rev.String()
}

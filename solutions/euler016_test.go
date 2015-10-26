package solutions

import "testing"
import (
	"math/big"
	"strconv"
)

/*
------------------------------------------------------------------------------------------------
 author: Eltjon Metko        ::         Power digit sum
------------------------------------------------------------------------------------------------
Problem URL : https://projecteuler.net/problem=16


2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.
What is the sum of the digits of the number 2^1000?
------------------------------------------------------------------------------------------------
*/

func Euler016() int {
	pow := new(big.Int).Exp(big.NewInt(2), big.NewInt(1000), nil).String()
	sum := 0
	for _, c := range pow {
		d, _ := strconv.Atoi(string(c))
		sum += d
	}
	return sum
}

func TestEuler016(t *testing.T) {
	expected := 1366
	got := Euler016()
	t.Logf("Answer: %v | Expected %v", got, expected)
	if got != expected {
		t.Fail()
	}
}

package solutions

import (
	"math/big"
	"testing"
)

/*
------------------------------------------------------------------------------------------------
 author: Eltjon Metko        ::         1000-digit Fibonacci number
------------------------------------------------------------------------------------------------
Problem URL : https://projecteuler.net/problem=25

The Fibonacci sequence is defined by the recurrence relation:
Fn = Fn−1 + Fn−2, where F1 = 1 and F2 = 1.
Hence the first 12 terms will be:
F1 = 1
F2 = 1
F3 = 2
F4 = 3
F5 = 5
F6 = 8
F7 = 13
F8 = 21
F9 = 34
F10 = 55
F11 = 89
F12 = 144
The 12th term, F12, is the first term to contain three digits.

What is the index of the first term in the Fibonacci sequence to contain 1000 digits?
------------------------------------------------------------------------------------------------
*/

func Euler025() int {

	//build 1e999 as string to initialise the max big.Int
	sl := "1"
	for i := 0; i < 999; i++ {
		sl += "0"
	}

	a, b := big.NewInt(0), big.NewInt(1)
	limit, _ := new(big.Int).SetString(sl, 10)
	i := 1
	for b.Cmp(limit) < 0 {
		a, b = b, new(big.Int).Add(a, b)
		i++
	}
	return i
}

func TestEuler025(t *testing.T) {
	expected := 4782
	got := Euler025()
	t.Logf("Answer: %v | Expected %v", got, expected)
	if got != expected {
		t.Fail()
	}
}

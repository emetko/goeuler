package solutions

import (
	"github.com/emetko/goeuler/utils"
	"math/big"
	"testing"
)

/*
------------------------------------------------------------------------------------------------
 author: Eltjon Metko        ::         Reciprocal cycles
------------------------------------------------------------------------------------------------
Problem URL : https://projecteuler.net/problem=26


A unit fraction contains 1 in the numerator. The decimal representation of the unit fractions with denominators 2 to 10 are given:

^1/2= 0.5
^1/3= 0.(3)
^1/4= 0.25
^1/5= 0.2
^1/6= 0.1(6)
^1/7= 0.(142857)
^1/8= 0.125
^1/9= 0.(1)
^1/10= 0.1

Where 0.1(6) means 0.166666..., and has a 1-digit recurring cycle. It can be seen that ^1/7 has a 6-digit recurring cycle.
Find the value of d < 1000 for which ^1/d contains the longest recurring cycle in its decimal fraction part.
------------------------------------------------------------------------------------------------
*/

func Euler026() int {
	maxL, maxD := 0, 0

	for d := 0; d < 1000; d++ {
		l := cycle(d)
		if l > maxL {
			maxL = l
			maxD = d
		}
	}
	return maxD
}

//10^(p−1) = 1 (mod p).
func cycle(d int) int {
	for i := 1; i <= d; i++ {
		bd := big.NewInt(int64(d))
		if 1 == new(big.Int).Mod(utils.Pow(10, i), bd).Int64() {
			return i
		}
	}
	return 0
}

func TestEuler026(t *testing.T) {
	expected := 983
	got := Euler026()
	t.Logf("Answer: %v | Expected %v", got, expected)
	if got != expected {
		t.Fail()
	}
}

package solutions

import (
	"github.com/emetko/goeuler/utils"
	"strconv"
	"testing"
)

/*
------------------------------------------------------------------------------------------------
 author: Eltjon Metko        ::         Digit cancelling fractions
------------------------------------------------------------------------------------------------
Problem URL : https://projecteuler.net/problem=33

The fraction ^49/98 is a curious fraction, as an inexperienced mathematician in attempting to simplify it
may incorrectly believe that ^49/98 = ^4/8, which is correct, is obtained by cancelling the 9s.
We shall consider fractions like, ^30/50 = ^3/5, to be trivial examples.
There are exactly four non-trivial examples of this type of fraction, less than one in value, and containing two digits in the numerator and denominator.
If the product of these four fractions is given in its lowest common terms, find the value of the denominator.
------------------------------------------------------------------------------------------------
*/

func Euler033() int {
	n, d := 1, 1
	for i := 11; i < 100; i++ {
		for j := i + 1; j < 100; j++ {
			if isFakeFraction(i, j) {
				n *= i
				d *= j
			}
		}
	}
	return d / utils.Gcd(n, d)
}

func isFakeFraction(a, b int) bool {
	if b%10 == 0 {
		return false
	}

	sa := strconv.Itoa(a)
	sb := strconv.Itoa(b)

	for i, ca := range sa {
		for j, cb := range sb {
			if ca == cb {
				na, _ := strconv.ParseFloat(sa[:i]+sa[i+1:], 64)
				nb, _ := strconv.ParseFloat(sb[:j]+sb[j+1:], 64)
				if na/nb == float64(a)/float64(b) {
					return true
				}
			}
		}
	}
	return false
}

func TestEuler033(t *testing.T) {
	expected := 100
	got := Euler033()
	t.Logf("Answer: %v | Expected %v", got, expected)
	if got != expected {
		t.Fail()
	}
}

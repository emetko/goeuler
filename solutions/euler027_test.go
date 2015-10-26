package solutions

import (
	"github.com/emetko/goeuler/utils"
	"testing"
)

/*
------------------------------------------------------------------------------------------------
 author: Eltjon Metko        ::         Quadratic primes
------------------------------------------------------------------------------------------------
Problem URL : https://projecteuler.net/problem=27

Euler discovered the remarkable quadratic formula:
n² + n + 41
It turns out that the formula will produce 40 primes for the consecutive values n = 0 to 39.
However, when n = 40, 40^2 + 40 + 41 = 40(40 + 1) + 41 is divisible by 41,
and certainly when n = 41, 41² + 41 + 41 is clearly divisible by 41.
The incredible formula  n² − 79n + 1601 was discovered, which produces 80 primes for the consecutive values n = 0 to 79.
The product of the coefficients, −79 and 1601, is −126479.

Considering quadratics of the form:    n² + an + b, where |a| < 1000 and |b| < 1000
where |n| is the modulus/absolute value of n            e.g. |11| = 11 and |−4| = 4

Find the product of the coefficients, a and b,
for the quadratic expression that produces the maximum number of primes for consecutive values of n,
starting with n = 0.
------------------------------------------------------------------------------------------------
*/

func Euler027(absLimit int) int {
	maxA, maxB, maxN := 0, 0, 0
	for a := -absLimit; a < absLimit; a++ {
		for b := -absLimit; b < absLimit; b++ {
			n := 0
			for {
				if utils.IsPrime(n*n + a*n + b) {
					n++
				} else {
					break
				}
			}
			if n > maxN {
				maxA, maxB, maxN = a, b, n
			}
		}
	}
	return maxA * maxB
}

func TestEuler027(t *testing.T) {
	expected := -59231
	got := Euler027(1000)
	t.Logf("Answer: %v | Expected %v", got, expected)
	if got != expected {
		t.Fail()
	}
}

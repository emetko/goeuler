package solutions

import (
	"github.com/emetko/goeuler/utils"
	"testing"
)

/*
------------------------------------------------------------------------------------------------
 author: Eltjon Metko        ::         Summation of primes
------------------------------------------------------------------------------------------------
Problem URL : https://projecteuler.net/problem=10

The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
Find the sum of all the primes below two million.
------------------------------------------------------------------------------------------------
*/

func Euler010() int {
	primes := utils.AtkinSieve(2000000)
	numPrimes := len(primes)
	sum := 0
	for i := 0; i < numPrimes; i++ {
		sum += primes[i]
	}
	return sum
}

func TestEuler010(t *testing.T) {
	expected := 142913828922
	got := Euler010()
	t.Logf("Answer: %v | Expected %v", got, expected)
	if got != expected {
		t.Fail()
	}
}

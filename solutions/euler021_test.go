package solutions

import (
	"github.com/emetko/goeuler/utils"
	"testing"
)

/*
------------------------------------------------------------------------------------------------
 author: Eltjon Metko        ::         Amicable numbers
------------------------------------------------------------------------------------------------
Problem URL : https://projecteuler.net/problem=21

Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).

If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.
For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.
Evaluate the sum of all the amicable numbers under 10000.
------------------------------------------------------------------------------------------------
*/

func Euler021(limit int) int {
	sum := 0
	candidates := make([]bool, limit)
	for i := 2; i < limit; i++ {
		if !candidates[i] {
			c := d(i)
			if c != i && d(c) == i {
				candidates[i] = true
				candidates[c] = true
			}
		}
	}

	for i := range candidates {
		if candidates[i] {
			sum += i
		}
	}
	return sum
}

func d(n int) (sum int) {
	dn := utils.Divisors(n) //Divisors returns also the number itself at the end of the slice
	// so we need to remove it later from loop
	dnl := len(dn)

	for _, v := range dn[:dnl-1] {
		sum += v
	}
	return sum
}

func TestEuler021(t *testing.T) {
	expected := 31626
	got := Euler021(10000)
	t.Logf("Answer: %v | Expected %v", got, expected)
	if got != expected {
		t.Fail()
	}
}

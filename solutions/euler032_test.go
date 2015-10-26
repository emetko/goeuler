package solutions

import (
	"github.com/emetko/goeuler/utils"
	"strconv"
	"testing"
)

/*
------------------------------------------------------------------------------------------------
 author: Eltjon Metko        ::         Pandigital products
------------------------------------------------------------------------------------------------
Problem URL : https://projecteuler.net/problem=32

We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once;
for example, the 5-digit number, 15234, is 1 through 5 pandigital.

The product 7254 is unusual, as the identity, 39 × 186 = 7254,
containing multiplicand, multiplier, and product is 1 through 9 pandigital.

Find the sum of all products whose multiplicand/multiplier/product identity can be written as a 1 through 9 pandigital.

HINT: Some products can be obtained in more than one way so be sure to only include it once in your sum.
------------------------------------------------------------------------------------------------
*/

func Euler032() int {
	prods := make(map[int]bool)
	for p := range utils.Perm(utils.Numerals[1:]) {
		_, _, prod := getIdentity(p)
		prods[prod] = true
	}
	sum := 0
	for k, v := range prods {
		if v {
			sum += k
		}
	}
	return sum
}

func getIdentity(sNum string) (int, int, int) {
	sl := len(sNum)

	for i := 0; i < sl/2+1; i++ {
		for j := i + 1; j < sl/2+2; j++ {
			a, _ := strconv.Atoi(sNum[:i])
			b, _ := strconv.Atoi(sNum[i:j])
			c, _ := strconv.Atoi(sNum[j:])
			if a*b == c {
				return a, b, c
			}
		}

	}
	return 0, 0, 0
}

func TestEuler032(t *testing.T) {
	expected := 45228
	got := Euler032()
	t.Logf("Answer: %v | Expected %v", got, expected)
	if got != expected {
		t.Fail()
	}
}

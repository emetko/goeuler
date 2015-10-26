package solutions

import (
	"github.com/emetko/goeuler/utils"
	"strconv"
	"testing"
)

/*
------------------------------------------------------------------------------------------------
 author: Eltjon Metko        ::         Factorial digit sum
------------------------------------------------------------------------------------------------
Problem URL : https://projecteuler.net/problem=20

n! means n × (n − 1) × ... × 3 × 2 × 1
For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.
Find the sum of the digits in the number 100!
------------------------------------------------------------------------------------------------
*/

func Euler020() int {
	sum := 0
	sf := utils.Fact(100).String()
	for _, c := range sf {
		d, _ := strconv.Atoi(string(c))
		sum += d
	}
	return sum
}

func TestEuler020(t *testing.T) {
	expected := 648
	got := Euler020()
	t.Logf("Answer: %v | Expected %v", got, expected)
	if got != expected {
		t.Fail()
	}
}

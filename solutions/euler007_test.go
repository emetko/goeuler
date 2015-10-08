package solutions

import (
	"testing"
	"github.com/emetko/goeuler/utils"
)

/*
------------------------------------------------------------------------------------------------
 author: Eltjon Metko        ::         10001st prime
------------------------------------------------------------------------------------------------
Problem URL : https://projecteuler.net/problem=7

By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
What is the 10 001st prime number?
------------------------------------------------------------------------------------------------
*/

func Euler007() int{
	return utils.NthPrime(10001)
}

func TestEuler007(t *testing.T){
	expected := 104743
	got := Euler007()
	t.Logf("Answer: %v | Expected %v",got, expected)
	if got!=expected{
		t.Fail()
	}
}

package solutions

import (
	"fmt"
	"github.com/emetko/goeuler/utils"
	"testing"
)

/*
------------------------------------------------------------------------------------------------
 author: Eltjon Metko        ::         Double-base palindromes
------------------------------------------------------------------------------------------------
Problem URL : https://projecteuler.net/problem=36


The decimal number, 585 = 1001001001 (binary), is palindromic in both bases.
Find the sum of all numbers, less than one million, which are palindromic in base 10 and base 2.
(Please note that the palindromic number, in either base, may not include leading zeros.)
------------------------------------------------------------------------------------------------
*/

func Euler036() int {
	sum := 0
	for i := 0; i < 1000000; i++ {
		if utils.IsPalindrome(fmt.Sprintf("%v", i)) && utils.IsPalindrome(fmt.Sprintf("%b", i)) {
			sum += i
		}
	}
	return sum
}

func TestEuler036(t *testing.T) {
	expected := 872187
	got := Euler036()
	t.Logf("Answer: %v | Expected %v", got, expected)
	if got != expected {
		t.Fail()
	}
}

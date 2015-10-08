package solutions

import "testing"
import (
	"github.com/emetko/goeuler/utils"
	"strconv"
)

/*
------------------------------------------------------------------------------------------------
 author: Eltjon Metko      ::          Largest palindrome product
------------------------------------------------------------------------------------------------
Problem URL : https://projecteuler.net/problem=4

A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
Find the largest palindrome made from the product of two 3-digit numbers.

------------------------------------------------------------------------------------------------
*/

func Euler004() int{
	max := 0
	for i:=100;i<1000;i++{
		for j:=100;j<1000;j++{
			prod := i*j
			if utils.IsPalindrome(strconv.Itoa(prod)) && prod > max{
				max = prod
			}
		}
	}
	return max
}

func TestEuler004(t *testing.T){
	expected := 906609
	got := Euler004()
	t.Logf("Answer: %v | Expected %v",got, expected)
	if got!=expected{
		t.Fail()
	}
}

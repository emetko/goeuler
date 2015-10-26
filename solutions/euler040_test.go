package solutions

import (
	"strconv"
	"testing"
)

/*
------------------------------------------------------------------------------------------------
 author: Eltjon Metko        ::         Champernowne&#39;s constant
------------------------------------------------------------------------------------------------
Problem URL : https://projecteuler.net/problem=40

An irrational decimal fraction is created by concatenating the positive integers:
0.123456789101112131415161718192021...
It can be seen that the 12^th digit of the fractional part is 1.
If dn represents the n^th digit of the fractional part, find the value of the following expression.
d1 × d10 × d100 × d1000 × d10000 × d100000 × d1000000
------------------------------------------------------------------------------------------------
*/

func Euler040() int {
	i, sentinel, product, stop := 1, 1, 1, 1000000

	currPos := 0
	for sentinel <= stop {
		s := strconv.Itoa(i)
		sl := len(s)
		currPos += sl
		if currPos >= sentinel {
			dPos := sl - (currPos - sentinel) - 1
			d, _ := strconv.Atoi(s[dPos : dPos+1])
			product *= d
			sentinel *= 10
		}
		i++
	}

	return product
}

func TestEuler040(t *testing.T) {
	expected := 210
	got := Euler040()
	t.Logf("Answer: %v | Expected %v", got, expected)
	if got != expected {
		t.Fail()
	}
}

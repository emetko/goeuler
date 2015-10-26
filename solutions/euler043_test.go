package solutions

import (
	"github.com/emetko/goeuler/utils"
	"strconv"
	"testing"
)

/*
------------------------------------------------------------------------------------------------
 author: Eltjon Metko        ::         Sub-string divisibility
------------------------------------------------------------------------------------------------
Problem URL : https://projecteuler.net/problem=43

The number, 1406357289, is a 0 to 9 pandigital number because it is made up of each of the digits 0 to 9 in some order, but it also has a rather interesting sub-string divisibility property.
Let d1 be the 1^st digit, d2 be the 2^nd digit, and so on. In this way, we note the following:
d2d3d4=406 is divisible by 2
d3d4d5=063 is divisible by 3
d4d5d6=635 is divisible by 5
d5d6d7=357 is divisible by 7
d6d7d8=572 is divisible by 11
d7d8d9=728 is divisible by 13
d8d9d10=289 is divisible by 17
Find the sum of all 0 to 9 pandigital numbers with this property.
------------------------------------------------------------------------------------------------
*/

/* solution:
since d4d5d6 must be divisible by 5: d6:=is in {0,5}
since d6d7d8 must be divisible by 11 d6 can not be 0 (that would leave only 011,022 etc that break the pandigital condition)
d4 must be even
will brute-force the rest
*/
func Euler043() int {
	s := 0
	for p := range utils.Perm(utils.Numerals) {
		if n, _ := strconv.Atoi(p[7:10]); n%17 != 0 {
			continue
		} else if n, _ := strconv.Atoi(p[6:9]); n%13 != 0 {
			continue
		} else if n, _ := strconv.Atoi(p[5:8]); n%11 != 0 {
			continue
		} else if n, _ := strconv.Atoi(p[4:7]); n%7 != 0 {
			continue
		} else if p[5:6] != "5" {
			continue
		} else if n, _ := strconv.Atoi(p[1:4]); n%2 != 0 {
			continue
		} else if n, _ := strconv.Atoi(p[2:5]); n%3 != 0 {
			continue
		}

		n, _ := strconv.Atoi(p)
		s += n
	}
	return s
}

func TestEuler043(t *testing.T) {
	expected := 16695334890
	got := Euler043()
	t.Logf("Answer: %v | Expected %v", got, expected)
	if got != expected {
		t.Fail()
	}
}

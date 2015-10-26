package solutions

import (
	"math"
	"strconv"
	"testing"
)

/*
------------------------------------------------------------------------------------------------
 author: Eltjon Metko        ::         Digit fifth powers
------------------------------------------------------------------------------------------------
Problem URL : https://projecteuler.net/problem=30

Surprisingly there are only three numbers that can be written as the sum of fourth powers of their digits:
1634 = 1^4 + 6^4 + 3^4 + 4^4
8208 = 8^4 + 2^4 + 0^4 + 8^4
9474 = 9^4 + 4^4 + 7^4 + 4^4
As 1 = 1^4 is not a sum it is not included.

The sum of these numbers is 1634 + 8208 + 9474 = 19316.
Find the sum of all the numbers that can be written as the sum of fifth powers of their digits.
------------------------------------------------------------------------------------------------
*/

/* solution:
let's find un upper limit to the numbers under consideration
4-digit numbers: max num would be 9^5+9^5+9^5+9^5 = 236196
5-digit numbers:                 9^5+9^5+9^5+9^5+9^5 = 295,245
6-digit numbers: max: 354294
7-digit dumbers: max: 413343

so we see that for 7 digits (and more) the sum of powers of 5 will have less digits than the numbers itself
in fact we need to consider up to the max 6 digit number 354294
*/
func Euler030(power float64) int {
	limit := int(6 * math.Pow(9.0, power))
	sum := 0
	for i := 2; i < limit; i++ {
		sNum := strconv.Itoa(i)
		iSum := 0
		for _, c := range sNum {
			ic, _ := strconv.Atoi(string(c))
			iSum += int(math.Pow(float64(ic), power))
		}

		if i == iSum {
			sum += i
		}
	}
	return sum
}

func TestEuler030(t *testing.T) {
	expected := 443839
	got := Euler030(5)
	t.Logf("Answer: %v | Expected %v", got, expected)
	if got != expected {
		t.Fail()
	}
}

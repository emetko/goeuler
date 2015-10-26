package solutions

import (
	"github.com/emetko/goeuler/utils"
	"strconv"
	"testing"
)

/*
------------------------------------------------------------------------------------------------
 author: Eltjon Metko        ::         Digit factorials
------------------------------------------------------------------------------------------------
Problem URL : https://projecteuler.net/problem=34


145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.
Find the sum of all numbers which are equal to the sum of the factorial of their digits.
Note: as 1! = 1 and 2! = 2 are not sums they are not included.
------------------------------------------------------------------------------------------------
*/

/*solution:
will use a similar analysis of that for problem 30.
we find that we need to consider numbers with max 7 digits as for 8-digit nums max sum of factorials will be with 7 digits.
so max = 9,999,999 at first pas.
but the max num we can obtain by summing factorials on a 7 digit number is 7*9! = 2,540,160
that will be our upper bound
*/
func Euler034() int {
	limit := int(7 * utils.Fact(9).Int64())

	facts := make([]int, 10)
	for i := 0; i < 10; i++ {
		facts[i] = int(utils.Fact(i).Int64())
	}

	sum := 0
	for i := 3; i < limit; i++ {
		sNum := strconv.Itoa(i)
		iSum := 0
		for _, c := range sNum {
			ic, _ := strconv.Atoi(string(c))
			iSum += facts[ic]
		}

		if i == iSum {
			sum += i
		}
	}
	return sum
}

func TestEuler034(t *testing.T) {
	expected := 40730
	got := Euler034()
	t.Logf("Answer: %v | Expected %v", got, expected)
	if got != expected {
		t.Fail()
	}
}

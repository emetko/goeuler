package solutions

import (
	"github.com/emetko/goeuler/utils"
	"testing"
)

/*
------------------------------------------------------------------------------------------------
 author: Eltjon Metko        ::         Number spiral diagonals
------------------------------------------------------------------------------------------------
Problem URL : https://projecteuler.net/problem=28

Starting with the number 1 and moving to the right in a clockwise direction a 5 by 5 spiral is formed as follows:
21 22 23 24 25
20  7  8  9 10
19  6  1  2 11
18  5  4  3 12
17 16 15 14 13
It can be verified that the sum of the numbers on the diagonals is 101.
What is the sum of the numbers on the diagonals in a 1001 by 1001 spiral formed in the same way?
------------------------------------------------------------------------------------------------
*/

/* solution:
			  73                      81
				 43	44 45 46 47 48 49
					21 22 23 24 25
					20  7  8  9 10
					19  6  1  2 11
					18  5  4  3 12
					17 16 15 14 13
				 37	36 35 34 33	32 31
			  65 					  57


layer 1: 					size 1
layer 2: 3,5,7,9			size 3  diff_from_topright: -2
layer 3: 13,17,21,25		size 5  diff_from_topright: -4
layer 4: 31,37,43,49		size 7  diff_from_topright: -6
layer 5: 57,65,73,81		size 9  diff_from_topright: -8

notice that top-right value in any layer is size^2
		and the other 3 corner values are obtained by repeatedly substracting size-1 from the top-right(anti-clockwise)
*/

func Euler028(size int) int {
	TRxDL := make([]int, 0, size-1)
	TLxDR := make([]int, 0, size-1)

	for i := 3; i <= size; i += 2 {
		tr := i * i
		TRxDL = append(TRxDL, tr, tr-2*(i-1))
		TLxDR = append(TLxDR, tr-(i-1), tr-3*(i-1))
	}

	return utils.Sum(TRxDL) + utils.Sum(TLxDR) + 1
}

func TestEuler028(t *testing.T) {
	expected := 669171001
	got := Euler028(1001)
	t.Logf("Answer: %v | Expected %v", got, expected)
	if got != expected {
		t.Fail()
	}
}

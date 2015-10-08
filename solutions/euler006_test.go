package solutions

import "testing"

/*
------------------------------------------------------------------------------------------------
 author: Eltjon Metko        ::         Sum square difference
------------------------------------------------------------------------------------------------
Problem URL : https://projecteuler.net/problem=6

The sum of the squares of the first ten natural numbers is,
1^2 + 2^2 + ... + 1^2 = 385
The square of the sum of the first ten natural numbers is,
(1 + 2 + ... + 10)^2 = 55^2 = 3025
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.
Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

------------------------------------------------------------------------------------------------
*/

func Euler006(max int) int{
	sumOfSquares,sum := 0, 0
	for i:=1;i<=max;i++{
		sumOfSquares += i*i
		sum += i
	}
	return sum*sum - sumOfSquares
}

func TestEuler006(t *testing.T){
	expected := 25164150
	got := Euler006(100)
	t.Logf("Answer: %v | Expected %v",got, expected)
	if got!=expected{
		t.Fail()
	}
}

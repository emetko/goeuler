package solutions

import (
	"testing"
	"github.com/emetko/goeuler/utils"
)

/*
------------------------------------------------------------------------------------------------
 author: Eltjon Metko        ::         Non-abundant sums
------------------------------------------------------------------------------------------------
Problem URL : https://projecteuler.net/problem=23

A perfect number is a number for which the sum of its proper divisors is exactly equal to the number.
For example, the sum of the proper divisors of 28 would be 1+2+4+7+14 = 28, which means that 28 is a perfect number.
A number n is called deficient if the sum of its proper divisors is less than n
	and it is called abundant if this sum exceeds n.

As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16,
the smallest number that can be written as the sum of two abundant numbers is 24.

By mathematical analysis,
it can be shown that all integers greater than 28123 can be written as the sum of two abundant numbers.
However, this upper limit cannot be reduced any further by analysis even though it is known that
the greatest number that cannot be expressed as the sum of two abundant numbers is less than this limit.

Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.
------------------------------------------------------------------------------------------------
*/

func Euler023(limit int) int{
	abNums := []int{}

	for i:=1;i<limit;i++{
		d := utils.Divisors(i)
		if utils.Sum(d[:len(d)-1]) > i{
			abNums = append(abNums,i)
		}
	}

	possibleSums := make(map[int]bool)

	for _,a1 := range abNums{
		for _,a2 := range abNums{
			if a1+a2<limit{
				possibleSums[a1+a2]=true
			}
		}
	}

	sum := 0
	for i:=0;i<limit;i++{
		if !possibleSums[i]{
			sum += i
		}
	}
	return sum
}

func TestEuler023(t *testing.T){
	expected := 4179871
	got := Euler023(28123)
	t.Logf("Answer: %v | Expected %v",got, expected)
	if got!=expected{
		t.Fail()
	}
}

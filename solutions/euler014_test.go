package solutions

import (
	"testing"
)

/*
------------------------------------------------------------------------------------------------
 author: Eltjon Metko        ::         Longest Collatz sequence
------------------------------------------------------------------------------------------------
Problem URL : https://projecteuler.net/problem=14

The following iterative sequence is defined for the set of positive integers:
n → n/2 (n is even)
n → 3n + 1 (n is odd)
Using the rule above and starting with 13, we generate the following sequence:
13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms.

Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.
Which starting number, under one million, produces the longest chain?
NOTE: Once the chain starts the terms are allowed to go above one million.
------------------------------------------------------------------------------------------------
*/

func Euler014() int{
	const limit = 1000000
	maxLen,maxNum := 1,1
	for i:= 1;i<limit;i++{
		len := 1
		for n:=i; n>1;len++{
			if n%2==0{
				n = n/2
			}else{
				n = 3*n+1
			}
		}
		if len>maxLen{
			maxLen = len
			maxNum = i
		}
	}
	return maxNum
}

func TestEuler014(t *testing.T){
	expected := 837799
	got := Euler014()
	t.Logf("Answer: %v | Expected %v",got, expected)
	if got!=expected{
		t.Fail()
	}
}

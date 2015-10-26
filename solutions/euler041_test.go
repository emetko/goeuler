package solutions

import (
	"github.com/emetko/goeuler/utils"
	"sort"
	"strconv"
	"testing"
)

/*
------------------------------------------------------------------------------------------------
 author: Eltjon Metko        ::         Pandigital prime
------------------------------------------------------------------------------------------------
Problem URL : https://projecteuler.net/problem=41


We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once.
For example, 2143 is a 4-digit pandigital and is also prime.
What is the largest n-digit pandigital prime that exists?
------------------------------------------------------------------------------------------------
*/

/* solution:
a 9-digit pandigital can not be prime becasue 1+2+3+4+5+6+7+8+9 = 45 thus it is divisible by 3
a 8-digit pandigital can not be prime becasue 1+2+3+4+5+6+7+8 = 36 thus it is divisible by 3
a 6-digit pandigital can not be prime becasue 1+2+3+4+5+6 = 21 thus it is divisible by 3
a 5-digit pandigital can not be prime becasue 1+2+3+4+5 = 15 thus it is divisible by 3
	will not consider pandigitals with less that 4 digits
	as we have at least a prime given with 4 digits that will be > that any of them
In conclussion we will iterate in decreasing order only on:
	7-digit and 4-digit odd pandigitals and stop as soon as we find a prime.
*/

func Euler041() int {
	//process 7 digit pandigitals
	p7d := make([]int, 0, 5040) //7!
	for v := range utils.Perm(utils.Numerals[1:8]) {
		n, _ := strconv.Atoi(v)
		p7d = append(p7d, n)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(p7d)))

	for i := 0; i < 5040; i++ {
		if utils.IsPrime(p7d[i]) {
			return p7d[i]
		}
	}

	//process 4 digit pandigitals
	p4d := make([]int, 0, 24) //4!
	for v := range utils.Perm(utils.Numerals[1:5]) {
		n, _ := strconv.Atoi(v)
		p4d = append(p4d, n)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(p4d)))

	for i := 0; i < 5040; i++ {
		if utils.IsPrime(p4d[i]) {
			return p4d[i]
		}
	}

	return -1
}

func TestEuler041(t *testing.T) {
	expected := 7652413
	got := Euler041()
	t.Logf("Answer: %v | Expected %v", got, expected)
	if got != expected {
		t.Fail()
	}
}

package solutions

import (
	"github.com/emetko/goeuler/utils"
	"github.com/jbarham/primegen.go"
	"testing"
)

/*
------------------------------------------------------------------------------------------------
 author: Eltjon Metko        ::         Distinct primes factors
------------------------------------------------------------------------------------------------
Problem URL : https://projecteuler.net/problem=47

The first two consecutive numbers to have two distinct prime factors are:
14 = 2 × 7
15 = 3 × 5
The first three consecutive numbers to have three distinct prime factors are:
644 = 2² × 7 × 23
645 = 3 × 5 × 43
646 = 2 × 17 × 19.
Find the first four consecutive integers to have four distinct prime factors. What is the first of these numbers?
------------------------------------------------------------------------------------------------
*/
var pCache = make([]uint64, 1000)

func Euler047() int {
	n, start := 647, 647
	pg := primegen.New()
	for i := 0; i < 1000; i++ {
		pCache[i] = pg.Next()
	}
	for {
		start = n
		found := 0
		for len(GetFactors(n)) == 4 {
			found++
			if found == 4 {
				return start
			}
			n++
		}
		n++
	}
	return 0
}

func GetFactors(n int) (fSlice []uint64) {
	if utils.IsPrime(n) {
		return []uint64{uint64(n)}
	}
	un := uint64(n)
	for _, v := range pCache {
		if un%v == 0 {
			fSlice = append(fSlice, v)
		}
	}

	return fSlice
}
func TestEuler047(t *testing.T) {
	expected := 134043
	got := Euler047()
	t.Logf("Answer: %v | Expected %v", got, expected)
	if got != expected {
		t.Fail()
	}
}

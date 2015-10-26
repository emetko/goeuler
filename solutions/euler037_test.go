package solutions

import (
	"testing"
	"strconv"
	"github.com/jbarham/primegen.go"
)

/*
------------------------------------------------------------------------------------------------
 author: Eltjon Metko        ::         Truncatable primes
------------------------------------------------------------------------------------------------
Problem URL : https://projecteuler.net/problem=37


The number 3797 has an interesting property.
Being prime itself, it is possible to continuously remove digits from left to right,
and remain prime at each stage: 3797, 797, 97, and 7. Similarly we can work from right to left: 3797, 379, 37, and 3.

Find the sum of the only eleven primes that are both truncatable from left to right and right to left.
NOTE: 2, 3, 5, and 7 are not considered to be truncatable primes.
------------------------------------------------------------------------------------------------
*/

func Euler037() int{
	np,sum := 0,0
	pg := primegen.New()

	cache := make(map[uint64]bool)
	cache[2] = true
	cache[3] = true
	cache[5] = true
	cache[7] = true
	pg.SkipTo(13)
	for np<11{
		n := pg.Next()
		cache[n]=true
		red := reductions(n)


		allP := true
		for _,r := range red{
			if ! cache[r]{
				allP = false
			}
		}
		if allP {
			np ++
			sum += int(n)
		}

	}

	return sum
}

func reductions(n uint64) []uint64{
	sn := strconv.FormatUint(n,10)
	nd := len(sn)
	r := make([]uint64,0,nd*2)
	for i:=1;i<nd;i++{
		c1,_ := strconv.Atoi(sn[i:])
		c2,_ := strconv.Atoi(sn[:nd-i])
		r = append(r,uint64(c1),uint64(c2))
	}
	return r
}

func TestEuler037(t *testing.T){
	expected := 748317
	got := Euler037()
	if got!=expected{
		t.Fail()
	}
}

package solutions

import (
	"fmt"
	"github.com/emetko/goeuler/utils"
	"github.com/jbarham/primegen.go"
	"strconv"
	"testing"
)

/*
------------------------------------------------------------------------------------------------
 author: Eltjon Metko        ::         Prime permutations
------------------------------------------------------------------------------------------------
Problem URL : https://projecteuler.net/problem=49

The arithmetic sequence, 1487, 4817, 8147, in which each of the terms increases by 3330, is unusual in two ways:
(i) each of the three terms are prime, and, (ii) each of the 4-digit numbers are permutations of one another.
There are no arithmetic sequences made up of three 1-, 2-, or 3-digit primes, exhibiting this property,
but there is one other 4-digit increasing sequence.
What 12-digit number do you form by concatenating the three terms in this sequence?
------------------------------------------------------------------------------------------------
*/

func Euler049() int {

	pg := primegen.New()
	pg.SkipTo(1489) //skip primes with less 4 digits and the example starting prime.
	for p := pg.Next(); p < 9999; p = pg.Next() {
		seq := make([]uint64, 0, 3)
		isPerm := GetPermutations(p)
		seq = append(seq, p)
		for n := p + 3330; utils.IsPrime(int(n)) && isPerm[n]; n += 3330 {
			seq = append(seq, n)
			if len(seq) == 3 {
				num, _ := strconv.Atoi(fmt.Sprintf("%v%v%v", seq[0], seq[1], seq[2]))
				return num
			}
		}
	}

	return 0
}

func GetPermutations(n uint64) map[uint64]bool {
	pMap := make(map[uint64]bool, 24)
	sp := utils.Perm(strconv.FormatUint(n, 10))
	for s := range sp {
		np, _ := strconv.ParseUint(s, 10, 64)
		pMap[np] = true
	}
	return pMap
}

func TestEuler049(t *testing.T) {
	expected := 296962999629
	got := Euler049()
	t.Logf("Answer: %v | Expected %v", got, expected)
	if got != expected {
		t.Fail()
	}
}

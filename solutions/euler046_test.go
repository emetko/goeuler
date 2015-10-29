package solutions

import (
	"github.com/jbarham/primegen.go"
	"math"
	"testing"
)

/*
------------------------------------------------------------------------------------------------
 author: Eltjon Metko        ::         Goldbach&#39;s other conjecture
------------------------------------------------------------------------------------------------
Problem URL : https://projecteuler.net/problem=46

It was proposed by Christian Goldbach that every odd composite number can be written as the sum of a prime and twice a square.
 9 =  7 + 2×1^2
15 =  7 + 2×2^2
21 =  3 + 2×3^2
25 =  7 + 2×3^2
27 = 19 + 2×2^2
33 = 31 + 2×1^2
It turns out that the conjecture was false.
What is the smallest odd composite that cannot be written as the sum of a prime and twice a square?
------------------------------------------------------------------------------------------------
*/

func Euler046() uint64 {
	isPrime := make(map[uint64]bool, 1000)

	pg := primegen.New()
	for i := 0; i < 1000; i++ {
		isPrime[pg.Next()] = true
	}

	found := false
	num := uint64(3)
	for !found {
		num += 2
		if isPrime[num] {
			num += 2
			continue
		}
		found = true //let's be optimists eh :)
		for k, _ := range isPrime {
			sq := math.Sqrt(float64((num - k) / 2))
			if sq == math.Floor(sq) {
				found = false
				break
			}
		}
	}
	return num
}

func TestEuler046(t *testing.T) {
	expected := uint64(5777)
	got := Euler046()
	t.Logf("Answer: %v | Expected %v", got, expected)
	if got != expected {
		t.Fail()
	}
}

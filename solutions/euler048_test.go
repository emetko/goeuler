package solutions

import (
	"math/big"
	"testing"
)

/*
------------------------------------------------------------------------------------------------
 author: Eltjon Metko        ::         Self powers
------------------------------------------------------------------------------------------------
Problem URL : https://projecteuler.net/problem=48

The series, 1^1 + 2^2 + 3^3 + ... + 10^10 = 10405071317.
Find the last ten digits of the series, 1^1 + 2^2 + 3^3 + ... + 1000^1000.
------------------------------------------------------------------------------------------------
*/

func Euler048() string {
	s := big.NewInt(0)
	p := big.NewInt(0)
	for i := int64(1); i < 1000; i++ {
		bi := big.NewInt(i)
		p.Exp(bi, bi, nil)
		s.Add(s, p)
	}
	ss := s.String()
	ssl := len(ss)
	return ss[ssl-10 : ssl]
}

func TestEuler048(t *testing.T) {
	expected := "9110846700"
	got := Euler048()
	t.Logf("Answer: %v | Expected %v", got, expected)
	if got != expected {
		t.Fail()
	}
}

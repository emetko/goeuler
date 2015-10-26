package solutions

import "testing"

/*
------------------------------------------------------------------------------------------------
 author: Eltjon Metko        ::         Integer right triangles
------------------------------------------------------------------------------------------------
Problem URL : https://projecteuler.net/problem=39

If p is the perimeter of a right angle triangle with integral length sides, {a,b,c}, there are exactly three solutions for p = 120.
{20,48,52}, {24,45,51}, {30,40,50}
For which value of p ≤ 1000, is the number of solutions maximised?
------------------------------------------------------------------------------------------------
*/

func Euler039() int {
	maxP, maxS := 0, 0
	//The perimeter is odd at the smollest one is 12 for (3,4,5)
	for p := 12; p <= 1000; p += 2 {
		s := 0
		for a := 1; a < p/3; a++ {
			for b := a + 1; b < (p-a)/2; b++ {
				c := p - a - b
				if a*a+b*b == c*c {
					s++
				}
			}
		}
		if s > maxS {
			maxP = p
			maxS = s
		}
	}
	return maxP
}

func TestEuler039(t *testing.T) {
	expected := 840
	got := Euler039()
	t.Logf("Answer: %v | Expected %v", got, expected)
	if got != expected {
		t.Fail()
	}
}

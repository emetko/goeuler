package solutions

import "testing"

/*
------------------------------------------------------------------------------------------------
 author: Eltjon Metko        ::         Special Pythagorean triplet
------------------------------------------------------------------------------------------------
Problem URL : https://projecteuler.net/problem=9

A Pythagorean triplet is a set of three natural numbers, a < b < c, for which a^2 + b^2 = c^2
For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.
There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
------------------------------------------------------------------------------------------------
*/

/*   loop optimisations:
	max(a) = sum/3 since a<b<c
	max(b) = (sum-a)/2 since b<c
	c is sum-a-b
*/
func Euler009(num int) int{
	for a:=1;a<num/3;a++{
		for b:=a+1;b<(num-a)/2;b++{
			c := num - a - b
			if c*c == a*a + b*b{
				return a*b*c
			}
		}
	}
	return 0
}

func TestEuler009(t *testing.T){
	expected := 31875000
	got := Euler009(1000)
	t.Logf("Answer: %v | Expected %v",got, expected)
	if got!=expected{
		t.Fail()
	}
}

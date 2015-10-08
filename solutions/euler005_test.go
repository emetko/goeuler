package solutions

import "testing"

/*
------------------------------------------------------------------------------------------------
 author: Eltjon Metko        ::         Smallest multiple
------------------------------------------------------------------------------------------------
Problem URL : https://projecteuler.net/problem=5

2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
------------------------------------------------------------------------------------------------
*/

/* Solution:
since every number n in 1..10 has at least a 2*n in 11...20
it's enough to check if X is divisable by y in 11...20

will start the loop at the known lcm(1..10) = 2520 and increment at that step
*/
func Euler005() int{

	for num := 2520;;num += 2520 {
		breakFound := false
		for i:=10;i<20;i++{
			if num%i!=0{
				breakFound = true
				break
			}
		}
		if !breakFound{
			return num
		}
	}
	return 0
}

func TestEuler005(t *testing.T){
	expected := 232792560
	got := Euler005()
	t.Logf("Answer: %v | Expected %v",got, expected)
	if got!=expected{
		t.Fail()
	}
}

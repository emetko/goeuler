/* Project Euler [https://projecteuler.net] Problem solutions */
package solutions

import "testing"

/*
------------------------------------------------------------------------------------------------
 author: Eltjon Metko        ::         Multiples of 3 and 5
------------------------------------------------------------------------------------------------
Problem URL : https://projecteuler.net/problem=1

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
------------------------------------------------------------------------------------------------
*/
func Euler001(max int) int{
	sum := 0
	for num:=0;num<max;num++{
		if num%3==0 || num%5==0{
			sum += num
		}
	}
	return sum
}

func TestEuler001(t *testing.T){
	expected := 233168
	got := Euler001(1000)
	t.Logf("Answer: %v | Expected %v",got, expected)
	if got!=expected {
		t.Fail()
	}
}

package solutions

import "testing"




/*
------------------------------------------------------------------------------------------------
 author: Eltjon Metko        ::         Largest prime factor
------------------------------------------------------------------------------------------------
Problem URL : https://projecteuler.net/problem=3

The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143 ?

------------------------------------------------------------------------------------------------
*/

func Euler003(num int) int{
	for i:=2;i<=num;i++ {
		for num % i == 0 {
			num = num / i
			if 1==num || i==num{
				return i
			}
		}
	}
	return num
}

func TestEuler003(t *testing.T){
	expected := 6857
	got := Euler003(600851475143)
	t.Logf("Answer: %v | Expected %v",got, expected)
	if got!=expected{
		t.Fail()
	}
}

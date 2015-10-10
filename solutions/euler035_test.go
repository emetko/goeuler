package solutions

import (
	"testing"
	"github.com/emetko/goeuler/utils"
	"fmt"
	"strconv"
)

/*
------------------------------------------------------------------------------------------------
 author: Eltjon Metko        ::         Circular primes
------------------------------------------------------------------------------------------------
Problem URL : https://projecteuler.net/problem=35

The number, 197, is called a circular prime because all rotations of the digits: 197, 971, and 719, are themselves prime.
There are thirteen such primes below 100: 2, 3, 5, 7, 11, 13, 17, 31, 37, 71, 73, 79, and 97.
How many circular primes are there below one million?
------------------------------------------------------------------------------------------------
*/

func Euler035() int{
	primes := utils.AtkinSieve(1000000)
	circulars := make(map[int]bool)
	for _,p:= range primes{
		if circulars[p]{
			continue
		}
		circular:=true
		rots := rotations(p)
		for _,r:= range rots{
			if !utils.IsPrime(r){
				circular = false
				break
			}
		}
		if circular{
			for _,r:= range rots{
				circulars[r]=true
			}
		}
	}

	fmt.Print(circulars)
	return len(circulars)
}

func rotations(num int) []int{
	sn := strconv.Itoa(num)
	sl := len(sn)
	rots := make([]int,sl)
	rots[0] = num
	for i:=0;i<sl-1;i++{
		sn = sn[sl-1:sl]+sn[:sl-1]
		rots[i+1],_= strconv.Atoi(sn)
	}
	return rots
}
func TestEuler035(t *testing.T){
	expected := 55
	got := Euler035()
	t.Logf("Answer: %v | Expected %v",got, expected)
	if got!=expected{
		t.Fail()
	}
}

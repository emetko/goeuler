package solutions

import (
	"testing"
	"github.com/emetko/goeuler/utils"
)

/*
------------------------------------------------------------------------------------------------
 author: Eltjon Metko        ::         Lexicographic permutations
------------------------------------------------------------------------------------------------
Problem URL : https://projecteuler.net/problem=24

A permutation is an ordered arrangement of objects.
For example, 3124 is one possible permutation of the digits 1, 2, 3 and 4.
If all of the permutations are listed numerically or alphabetically, we call it lexicographic order.

The lexicographic permutations of 0, 1 and 2 are:    012   021   102   120   201   210

What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?
------------------------------------------------------------------------------------------------
*/

func Euler024(pNum int) string{
	i:= 1
	for v := range Perm(utils.Numerals){
		if i == pNum{
			return v
		}
		i++
	}
	return ""
}

func Perm(iter string) <-chan string {
	c := make(chan string)
	go func(c chan string) {
		defer close(c)
		permIter(c, "", iter)
	}(c)
	
	return c
}


func permIter(c chan string, combo string, alphabet string) {
	if len(alphabet) == 0 {
		c <- combo
		return
	}

	for i, ch := range alphabet {
		newCombo := combo + string(ch)
		newAlphabet := alphabet[:i]+alphabet[i+1:]
		permIter(c, newCombo, newAlphabet)
	}

}

func TestEuler024(t *testing.T){
	expected := "2783915460"
	got := Euler024(1000000)
	t.Logf("Answer: %v | Expected %v",got, expected)
	if got!=expected{
		t.Fail()
	}
}

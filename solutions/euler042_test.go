package solutions

import (
	"github.com/emetko/goeuler/utils"
	"io/ioutil"
	"strings"
	"testing"
)

/*
------------------------------------------------------------------------------------------------
 author: Eltjon Metko        ::         Coded triangle numbers
------------------------------------------------------------------------------------------------
Problem URL : https://projecteuler.net/problem=42

The n^th term of the sequence of triangle numbers is given by, tn = ½n(n+1); so the first ten triangle numbers are:
1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...
By converting each letter in a word to a number corresponding to its alphabetical position and adding these values we form a word value.
For example, the word value for SKY is 19 + 11 + 25 = 55 = t10.
If the word value is a triangle number then we shall call the word a triangle word.
Using words.txt (right click and "Save Link/Target As..."), a 16K text file containing nearly two-thousand common English words,
how many are triangle words?
------------------------------------------------------------------------------------------------
*/

func Euler042() int {
	ns, err := ioutil.ReadFile("../resources/p042_words.txt")
	if err != nil {
		panic("could not read words file")
	}

	words := strings.Split(strings.Trim(string(ns), "\""), "\",\"")

	n := 0
	for _, w := range words {
		if utils.IsTriangular(utils.Score(w)) {
			n++
		}
	}
	return n
}

func TestEuler042(t *testing.T) {
	expected := 162
	got := Euler042()
	t.Logf("Answer: %v | Expected %v", got, expected)
	if got != expected {
		t.Fail()
	}
}

package solutions

import (
	"strings"
	"testing"
)

/*
------------------------------------------------------------------------------------------------
 author: Eltjon Metko        ::         Number letter counts
------------------------------------------------------------------------------------------------
Problem URL : https://projecteuler.net/problem=17

If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.
If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?

NOTE: Do not count spaces or hyphens.
For example, 342 (three hundred and forty-two) contains 23 letters
         and 115 (one hundred and fifteen) contains 20 letters.
         The use of "and" when writing out numbers is in compliance with British usage.
------------------------------------------------------------------------------------------------
*/

func Euler017(limit int) int {
	cc := 0
	for i := 0; i <= limit; i++ {
		s := word(i)
		cc += len(s) - strings.Count(s, " ") - strings.Count(s, "-")
	}
	return cc
}

func word(num int) string {
	if num == 1000 {
		return "one thousand"
	}
	words := []string{
		"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
		"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
	tens := []string{"twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

	units := words[0:10]
	for _, t := range tens {
		n := ""
		for _, u := range units {
			if u == "" {
				n = t
			} else {
				n = t + "-" + u
			}
			words = append(words, n)
		}
	}

	hundreds := num / 100
	rest := num % 100

	w := ""
	if hundreds > 0 {
		w += words[hundreds] + " hundred"
		if rest > 0 {
			w += " and "
		}
	}
	w += words[rest]

	return w
}

func TestEuler017(t *testing.T) {
	expected := 21124
	got := Euler017(1000)
	t.Logf("Answer: %v | Expected %v", got, expected)
	if got != expected {
		t.Fail()
	}
}

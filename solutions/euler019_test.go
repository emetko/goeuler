package solutions

import (
	"testing"
	"time"
)

/*
------------------------------------------------------------------------------------------------
 author: Eltjon Metko        ::         Counting Sundays
------------------------------------------------------------------------------------------------
Problem URL : https://projecteuler.net/problem=19


You are given the following information, but you may prefer to do some research for yourself.
1 Jan 1900 was a Monday.
Thirty days has September,
April, June and November.
All the rest have thirty-one,
Saving February alone,
Which has twenty-eight, rain or shine.
And on leap years, twenty-nine.

A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.

How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?
------------------------------------------------------------------------------------------------
*/

func Euler019() int {
	num := 0
	for y := 1901; y < 2001; y++ {
		for m := 1; m < 13; m++ {
			d := time.Date(y, time.Month(m), 1, 15, 0, 0, 0, time.UTC)
			if d.Weekday() == time.Sunday {
				num++
			}
		}
	}
	return num
}

func TestEuler019(t *testing.T) {
	expected := 171
	got := Euler019()
	t.Logf("Answer: %v | Expected %v", got, expected)
	if got != expected {
		t.Fail()
	}
}

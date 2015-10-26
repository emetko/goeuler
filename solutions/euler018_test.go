package solutions

import (
	"testing"
)

/*
------------------------------------------------------------------------------------------------
 author: Eltjon Metko        ::         Maximum path sum I
------------------------------------------------------------------------------------------------
Problem URL : https://projecteuler.net/problem=18

By starting at the top of the triangle below and moving to adjacent numbers on the row below, the maximum total from top to bottom is 23.
3
7 4

2 4 6

8 5 9 3
That is, 3 + 7 + 4 + 9 = 23.
Find the maximum total from top to bottom of the triangle below:
75

95 64

17 47 82

18 35 87 10

20 04 82 47 65

19 01 23 75 03 34

88 02 77 73 07 63 67

99 65 04 28 06 16 70 92

41 41 26 56 83 40 80 70 33

41 48 72 33 47 32 37 16 94 29

53 71 44 65 25 43 91 52 97 51 14

70 11 33 28 77 73 17 78 39 68 17 57

91 71 52 38 17 14 91 43 58 50 27 29 48

63 66 04 68 89 53 67 30 73 16 69 87 40 31

04 62 98 27 23 09 70 98 73 93 38 53 60 04 23
NOTE: As there are only 16384 routes, it is possible to solve this problem by trying every route. However, Problem 67, is the same challenge with a triangle containing one-hundred rows; it cannot be solved by brute force, and requires a clever method! ;o)
------------------------------------------------------------------------------------------------
*/
var test = [][]int{
	{3},
	{7, 4},
	{2, 4, 6},
	{8, 5, 9, 3}}
var data = [][]int{
	{75},
	{95, 64},
	{17, 47, 82},
	{18, 35, 87, 10},
	{20, 04, 82, 47, 65},
	{19, 01, 23, 75, 03, 34},
	{88, 02, 77, 73, 07, 63, 67},
	{99, 65, 04, 28, 06, 16, 70, 92},
	{41, 41, 26, 56, 83, 40, 80, 70, 33},
	{41, 48, 72, 33, 47, 32, 37, 16, 94, 29},
	{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14},
	{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57},
	{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48},
	{63, 66, 04, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31},
	{04, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 04, 23}}

/* solution:
	start from the triangle base (-1 level)
	walking up, substitute each value with the max sum possible
    max sum for the whole triangle will be at 0,0
*/

func Euler018(triang [][]int) int {
	depth := len(triang)
	for i := depth - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			triang[i][j] = triang[i][j] + max(triang[i+1][j], triang[i+1][j+1])
		}
	}
	return triang[0][0]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func TestEuler018(t *testing.T) {
	expected := 1074
	got := Euler018(data)
	t.Logf("Answer: %v | Expected %v", got, expected)
	if got != expected {
		t.Fail()
	}
}

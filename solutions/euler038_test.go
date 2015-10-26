package solutions

import (
	"testing"
	"fmt"
	"strconv"
	"github.com/emetko/goeuler/utils"
)

/*
------------------------------------------------------------------------------------------------
 author: Eltjon Metko        ::         Pandigital multiples
------------------------------------------------------------------------------------------------
Problem URL : https://projecteuler.net/problem=38

Take the number 192 and multiply it by each of 1, 2, and 3:
192 × 1 = 192
192 × 2 = 384
192 × 3 = 576
By concatenating each product we get the 1 to 9 pandigital, 192384576. We will call 192384576 the concatenated product of 192 and (1,2,3)
The same can be achieved by starting with 9 and multiplying by 1, 2, 3, 4, and 5, giving the pandigital, 918273645, which is the concatenated product of 9 and (1,2,3,4,5).
What is the largest 1 to 9 pandigital 9-digit number that can be formed as the concatenated product of an integer with (1,2, ... , n) where n > 1?
------------------------------------------------------------------------------------------------
*/

/* solution:
Since n>1 max_num will have no more than 4 digits so first upper limit is 9876

Since we have one example 918273645, our number must start with 9 to be greater

any m-digit number starting with 9 when multiplied by 2,3 etc gives m+1 digits as result
for m = 1 we already have the result given so we must exclude it
for m = 2 we have either 2+3+3<9 digits for (n=3) or  2+3+3+3>9 for n=4
for m = 3 we have either 3+4<9 digits for n=2 or 3+4+4>9 digits for n=3
for m = 4 we have 4+5=9 digits for n=2 so first lower limit is 9123

*/

func Euler038() int{
	m := 0
	for n:=9123;n<=9876;n++{
		sn := fmt.Sprintf("%v%v",n,n*2)
		if utils.IsPandigital(sn){
			p,_:=strconv.Atoi(sn)
			m = max(m,p)
		}
	}
	return m
}


func TestEuler038(t *testing.T){
	expected := 932718654
	got := Euler038()
	t.Logf("Answer: %v | Expected %v",got, expected)
	if got!=expected{
		t.Fail()
	}
}

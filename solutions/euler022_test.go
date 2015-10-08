package solutions

import (
	"testing"
	"io/ioutil"
	"strings"
	"sort"
	"github.com/emetko/goeuler/utils"
)

/*
------------------------------------------------------------------------------------------------
 author: Eltjon Metko        ::         Names scores
------------------------------------------------------------------------------------------------
Problem URL : https://projecteuler.net/problem=22

Using names.txt (right click and "Save Link/Target As..."),
a 46K text file containing over five-thousand first names,
begin by sorting it into alphabetical order.
Then working out the alphabetical value for each name,
multiply this value by its alphabetical position in the list to obtain a name score.
For example, when the list is sorted into alphabetical order,
COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53,
is the 938th name in the list. So, COLIN would obtain a score of 938 × 53 = 49714.

What is the total of all the name scores in the file?
------------------------------------------------------------------------------------------------
*/

func Euler022() int{

	ns,err := ioutil.ReadFile("../resources/p022_names.txt")
	if err != nil{
		panic("could not read names file")
	}

	names := strings.Split(strings.Trim(string(ns),"\""), "\",\"")
	sort.Strings(names)
	total := 0
	for i,v := range names{
		total += score(v)*(i+1)
	}

	return total
}

func score(str string) (sc int){
	scMap := make(map[rune]int,26)
	for i,r:=range strings.ToUpper(utils.Alphabet){
		scMap[r]=i+1
	}
	for _,r := range str{
		sc += scMap[r]
	}
	return sc
}

func TestEuler022(t *testing.T){
	expected := 871198282
	got := Euler022()
	t.Logf("Answer: %v | Expected %v",got, expected)
	if got!=expected{
		t.Fail()
	}
}

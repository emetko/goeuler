package utils

import (
	"sort"
	"strings"
)

type SortRunes []rune

func (s SortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s SortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(SortRunes(r))
	return string(r)
}

func IsPandigital(s string) bool {
	rs := []rune(s)
	sort.Sort(SortRunes(rs))
	return string(rs) == Numerals[1:len(s)+1]
}

func Score(str string) (sc int) {
	scMap := make(map[rune]int, 26)
	for i, r := range strings.ToUpper(Alphabet) {
		scMap[r] = i + 1
	}
	for _, r := range str {
		sc += scMap[r]
	}
	return sc
}

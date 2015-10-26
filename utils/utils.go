package utils

const Alphabet = "abcdefghijklmnopqrstuvwxyz"
const Numerals = "0123456789"

func  nthdigit(num, pos int) int{
	powersOf10 := []int{1, 10, 100, 1000, 10000}
	return ((num / powersOf10[pos]) % 10)
}

func IsPalindrome(str string) bool{
	sl := len(str)
	for i:=0;i<sl/2;i++{
		if str[i]!=str[sl-i-1]{
			return false
		}
	}
	return true
}

func Max(is []int) int{
	maxVal := 0;
	for _,v := range is{
		if v>maxVal{
			maxVal = v
		}
	}
	return maxVal
}

func Sum(is []int) int{
	sum := 0
		for _,v :=range is{
			sum += v
		}
	return sum
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


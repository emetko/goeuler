package utils

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

package utils

import (
	"math"
	"math/big"
)


func IsPrime(n int) bool{
	if n == 2 || n == 3 {
		return true
	}
	if n < 2 || n % 2 == 0 || n % 3 == 0 {
		return false
	}
	if n < 9 {
		return true}

	r := math.Sqrt(float64(n))
	i := 5
	for float64(i) <= r {
		if n % i == 0 || n % (i + 2) == 0 {
			return false
		}
		i += 6
	}
	return true
}

func NthPrime(n int) int{
	if n==1{
		return 2
	}
	num := 3
	for count:=2;count<=n;num+=2{
		if IsPrime(num){
			if count==n{
				return num
			}else {
				count++
			}
		}
	}
	return num
}

func AtkinSieve(maxNum int) []int {
	var x, y, n int
	nsqrt := math.Sqrt(float64(maxNum))

	is_prime := make([]bool, maxNum)

	//consider only up to sqrt(N)
	for x = 1; float64(x) <= nsqrt; x++ {
		for y = 1; float64(y) <= nsqrt; y++ {
			n = 4 * (x * x) + y * y
			if n <= maxNum && (n % 12 == 1 || n % 12 == 5) {
				is_prime[n] = !is_prime[n]
			}
			n = 3 * (x * x) + y * y
			if n <= maxNum && n % 12 == 7 {
				is_prime[n] = !is_prime[n]
			}
			n = 3 * (x * x) - y * y
			if x > y && n <= maxNum && n % 12 == 11 {
				is_prime[n] = !is_prime[n]
			}
		}
	}

	//filter out multiples of primes
	for n = 5; float64(n) <= nsqrt; n++ {
		if is_prime[n] {
			for y = n * n; y < maxNum; y += n * n {
				is_prime[y] = false
			}
		}
	}
	is_prime[2] = true
	is_prime[3] = true

	primes := make([]int, 0, maxNum)
	for x = 0; x < len(is_prime) - 1; x++ {
		if is_prime[x] {
			primes = append(primes, x)
		}
	}
	return primes
}

func GetFactors(num int) (fSlice []int) {

	primes := AtkinSieve(num/2)

	fSlice = append(fSlice,1)
	for _,v:=range primes {
		if num % v == 0 {
			fSlice = append(fSlice,v)
		}
	}
	fSlice = append(fSlice,num)

	return fSlice
}

func Divisors(num int) (dSlice []int){
	sqrtN := int(math.Sqrt(float64(num)))

	dSlice = append(dSlice,1)
	i:=2
	for ;i<=sqrtN;i++ {
		if num%i==0{
			dSlice = append(dSlice,i)
			if i*i!=num {
				dSlice = append(dSlice,num/i)
			}
		}
	}

	dSlice = append(dSlice,num)
	return dSlice
}

func NthTriangleNum(n int) (sum int){
	for i:=0;i<n;i++{
		sum += i
	}
	return sum
}

func Fact(n int)  *big.Int{
	return new(big.Int).MulRange(1,int64(n))
}

func Comb(n int,k int) uint64{
	return new(big.Int).Div(Fact(n),new(big.Int).Mul(Fact(k),Fact(n-k))).Uint64()
}

func Pow(b int,e int) *big.Int{
	bb := big.NewInt(int64(b))
	if e==1{
		return bb
	}

	res := big.NewInt(1)
	q,r := e/2,e%2
	ires := Pow(b,q)
	res.Mul(ires,ires)
	if r>0{
		res.Mul(res,bb)
	}
	return res
}
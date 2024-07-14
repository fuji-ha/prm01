package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("数字nを入力してください: ")
	fmt.Scan(&n)

	primes := generatePrimes(n)
	fmt.Printf("%d個の素数: %v\n", n, primes)
}

// n個の素数を生成する関数
func generatePrimes(n int) []int {
	primes := []int{}
	count := 0
	num := 2

	for count < n {
		if isPrime(num) {
			primes = append(primes, num)
			count++
		}
		num++
	}

	return primes
}

// 素数かどうかを判定する関数
func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	sqrt := int(math.Sqrt(float64(num)))
	for i := 2; i <= sqrt; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

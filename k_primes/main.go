package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	sq := int(math.Sqrt(float64(n)))
	for i := 2; i <= sq; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func countFactors(n, result int) int {
	for i := 2; i < n; i++ {
		if isPrime(i) {
			if n%i == 0 {
				result++
				return countFactors(n/i, result)
			}
		}
	}
	return result
}

func CountKprimes(g, m, n int) []int {
	var arr []int
	for i := m; i <= n; i++ {
		result := 1
		if countFactors(i, result) == g {
			arr = append(arr, i)
		}
	}
	return arr
}

func main() {
	fmt.Println(CountKprimes(5, 500, 600))        // 500, 520, 552, 567, 588, 592, 594
	fmt.Println(CountKprimes(5, 1000, 1100))      // 1020, 1026, 1032, 1044, 1050, 1053, 1064,   1072, 1092, 1100
	fmt.Println(CountKprimes(12, 100000, 100100)) // nil
}

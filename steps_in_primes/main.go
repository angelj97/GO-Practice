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

func Step(g, m, n int) []int {
	if m == n {
		return nil
	}
	if isPrime(m) {
		if isPrime(m + g) {
			arr := []int{0, 0}
			arr[0] = m
			arr[1] = m + g
			return arr
		}
	}
	arr := Step(g, m+1, n)
	return arr
}

func main() {
	fmt.Println(Step(2, 100, 110))       // 101, 103
	fmt.Println(Step(4, 100, 110))       // 103, 107
	fmt.Println(Step(6, 100, 110))       // 101, 107
	fmt.Println(Step(8, 300, 400))       // 359, 367
	fmt.Println(Step(10, 300, 400))      // 307, 317
	fmt.Println(Step(11, 30000, 100000)) // nil
}

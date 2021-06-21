package main

import (
	"fmt"
)

var prev map[int]int

//func fib(n int) int{
//	if n <= 2 {
//		return 1
//	}
//	return fib(n-1) + fib(n-2)
//}

//func fib(a, b, n int) int{
//	if n <= 2 {
//		return b
//	}
//	temp := a + b
//	a = b
//	b = temp
//	return fib(a,b,n-1)
//}

func fib(n int) int {
	if n <= 2 {
		return 1
	}
	if prev[n] != 0 {
		return prev[n]
	} else {
		prev[n] = fib(n-1) + fib(n-2)
	}
	return prev[n]
}

func main() {
	prev = make(map[int]int)
	prev[1] = 1
	prev[2] = 1
	fmt.Println(fib(6))
	fmt.Println(fib(7))
	fmt.Println(fib(8))
	fmt.Println(fib(50))
	fmt.Println(fib(3))
	fmt.Println(fib(83))
	fmt.Println(fib(92))
}

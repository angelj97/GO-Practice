package main

import "fmt"

func canSumMap(target int, memo map[int]bool, numbers ...int) bool {
	if _, exists := memo[target]; exists {
		return memo[target]
	}
	if target == 0 {
		return true
	}
	if target < 0 {
		return false
	}
	for _, number := range numbers {
		remainder := target - number
		if canSumMap(remainder, memo, numbers...) {
			memo[target] = true
			return true
		}
	}
	memo[target] = false
	return false
}

func canSum(target int, numbers ...int) bool {
	memo := make(map[int]bool)
	return canSumMap(target, memo, numbers...)
}

func main() {
	fmt.Println(canSum(7, 2, 3))
	fmt.Println(canSum(7, 5, 3, 4, 7))
	fmt.Println(canSum(7, 2, 4))
	fmt.Println(canSum(8, 2, 3, 5))
	fmt.Println(canSum(300, 7, 14))
}

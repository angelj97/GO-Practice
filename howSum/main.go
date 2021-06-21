package main

import "fmt"

func howSumMap(target int, memo map[int][]int, numbers ...int) []int {
	if _, exists := memo[target]; exists {
		return memo[target]
	}
	if target == 0 {
		return []int{}
	}
	if target < 0 {
		return nil
	}
	for _, num := range numbers {
		remainder := target - num
		var result []int = howSumMap(remainder, memo, numbers...)
		if result != nil {
			memo[target] = append(result, num)
			return memo[target]
		}
	}
	memo[target] = nil
	return nil
}

func howSum(target int, numbers ...int) []int {
	memo := make(map[int][]int)
	return howSumMap(target, memo, numbers...)
}

func main() {
	fmt.Println(howSum(7, 2, 3))
	fmt.Println(howSum(7, 5, 3, 4, 7))
	fmt.Println(howSum(7, 2, 4))
	fmt.Println(howSum(8, 2, 3, 5))
	fmt.Println(howSum(300, 7, 14))
}

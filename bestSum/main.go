package main

import "fmt"

func bestSumMap(target int, memo map[int][]int, numbers ...int) []int {
	if _, exists := memo[target]; exists {
		return memo[target]
	}
	if target == 0 {
		return []int{}
	}
	if target < 0 {
		return nil
	}
	var shortest []int = nil
	for _, num := range numbers {
		remainder := target - num
		var result []int = bestSumMap(remainder, memo, numbers...)
		if result != nil {
			combination := append(result, num)
			if shortest == nil || len(combination) < len(shortest) {
				shortest = combination
			}
		}
	}
	memo[target] = shortest
	return shortest
}

func bestSum(target int, numbers ...int) []int {
	memo := make(map[int][]int)
	return bestSumMap(target, memo, numbers...)
}

func main() {
	fmt.Println(bestSum(7, 5, 3, 4, 7))
	fmt.Println(bestSum(8, 2, 3, 5))
	fmt.Println(bestSum(8, 1, 4, 5))
	fmt.Println(bestSum(100, 1, 2, 5, 25))
}

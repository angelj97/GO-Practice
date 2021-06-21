package main

import "fmt"

// func sumArr(arr []int, n, sum int) int {
// 	if n == 0 {
// 		return sum
// 	}
// 	sum += arr[n-1]
// 	n--
// 	return sumArr(arr, n, sum)
// }

// func sumArray(arr []int) int {
// 	sum := 0
// 	n := len(arr)
// 	return sumArr(arr, n, sum)
// }

func sumArray(arr []int) int {
	sum := 0
	for _, number := range arr {
		sum += number
	}
	return sum
}

func main() {
	array := []int{5, 3, 5, 3, 2, 1, 6, 3, 3, 13}
	fmt.Println(sumArray(array))
}

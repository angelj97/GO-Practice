package kata

import (
	"fmt"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	// Split numbers in the string into an array
	arr := strings.Fields(in)
	// Assign highest and lowest number to a variable as the base
	// Since the array is still full of strings we convert first convert it to int
	high, _ := strconv.Atoi(arr[0])
	low, _ := strconv.Atoi(arr[0])
	// Loop over the array we created
	for _, numS := range arr {
		num, _ := strconv.Atoi(numS)
		// Assign new highest and lowest
		if num >= high {
			high = num
		} else if num <= low {
			low = num
		}
	}
	// Concatinate high and low into a string
	res := (fmt.Sprint(high) + " " + fmt.Sprint(low))
	return res
}

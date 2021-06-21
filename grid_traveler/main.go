package main

import "fmt"

var gridMap map[string]int
var key string

func gridTraveler(x, y int) int {
	key = fmt.Sprint(x) + "," + fmt.Sprint(y)
	if gridMap[key] != 0 {
		return gridMap[key]
	}
	if x == 1 && y == 1 {
		return 1
	}
	if x == 0 || y == 0 {
		return 0
	}
	gridMap[key] = gridTraveler(x-1, y) + gridTraveler(x, y-1)
	return gridMap[key]
}

func main() {
	gridMap = make(map[string]int)
	fmt.Println(gridTraveler(1, 1))
	fmt.Println(gridTraveler(2, 3))
	fmt.Println(gridTraveler(3, 2))
	fmt.Println(gridTraveler(2, 3))
	fmt.Println(gridTraveler(3, 3))
	fmt.Println(gridTraveler(18, 18))
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	set := bufio.NewReader(os.Stdin)
	setChars, _ := set.ReadString()
	var uChars []rune
	var count int
	for i, char := range setChars {
		switch char {
		case uChars:
			count++
		}
	}
	fmt.Println(len(uChars))
}

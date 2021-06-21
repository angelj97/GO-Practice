package main

import (
	"fmt"
	"strings"
)

func canConstruct(target string, wordBank ...string) bool {
	memo := make(map[string]bool)
	return canConstructMemo(target, memo, wordBank...)
}

func canConstructMemo(target string, memo map[string]bool, wordBank ...string) bool {
	if _, exists := memo[target]; exists {
		return memo[target]
	}
	if target == "" {
		return true
	}
	for _, word := range wordBank {
		if strings.HasPrefix(target, word) {
			newTarget := strings.TrimPrefix(target, word)
			if canConstructMemo(newTarget, memo, wordBank...) {
				memo[target] = true
				return true
			}
		}
	}
	memo[target] = false
	return false
}

func main() {
	fmt.Println(canConstruct("abcdef", "ab", "abc", "cd", "def", "abcd"))
	fmt.Println(canConstruct("skateboard", "bo", "rd", "ate", "t", "ska", "sk", "boar"))
	fmt.Println(canConstruct("enterapotentpot", "a", "p", "ent", "enter", "ot", "o", "t"))
	fmt.Println(canConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", "e", "ee", "eee", "eeee", "eeeee", "eeeeee"))
}

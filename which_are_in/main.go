package kata

import (
	"sort"
	"strings"
)

func InArray(array1 []string, array2 []string) []string {
	// Have to initialize slice with 0 to return [] instead of nil
	result := make([]string, 0)
	// Created map to keep track of any duplicates
	added := make(map[string]bool)
	for _, a1 := range array1 {
		// Checks of duplicates, skips if it already exists
		if _, exists := added[a1]; exists {
			continue
		}
		// Adds any new results into the map and array
		for _, a2 := range array2 {
			if strings.Contains(a2, a1) {
				result = append(result, a1)
				added[a1] = true
				break
			}
		}
	}
	sort.Strings(result)
	return result
}

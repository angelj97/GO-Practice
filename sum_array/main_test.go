package main

import "testing"
import "reflect"

func TestTableSumArray(t *testing.T) {
	got := sumArray([]int{1, 1}, []int{1, 2, 3, 4, 5}, []int{3, -3})
	want := []int{2, 15, 0}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d, want %d", got, want)
	}
}

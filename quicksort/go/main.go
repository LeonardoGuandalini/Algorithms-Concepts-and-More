package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Implementation of QuickSort in Go
func main() {
	var array []int
	for i := 0; i < 100000; i++ {
		array = append(array, rand.Intn(100))
	}
	beggining := time.Now()

	(QuickSort(array))
	end := time.Now()
	fmt.Println(end.Sub(beggining))
}

func QuickSort(sequence []int) []int {
	size := len(sequence)

	if size <= 1 {
		return sequence
	}

	pivot, left := Pop(sequence)

	var greater []int
	var smaller []int

	for _, item := range left {

		if item > pivot {
			greater = append(greater, item)
		} else {
			smaller = append(smaller, item)
		}
	}
	pivot_list := []int{pivot}
	var sorted_values []int
	sorted_values = append(sorted_values, QuickSort(smaller)...)
	sorted_values = append(sorted_values, pivot_list...)
	sorted_values = append(sorted_values, QuickSort(greater)...)

	return sorted_values
}
func Pop(sequence []int) (int, []int) {

	size := len(sequence)
	if size < 1 {
		return sequence[0], nil
	} else {
		element := sequence[size-1]
		return element, sequence[:size-1]
	}
}

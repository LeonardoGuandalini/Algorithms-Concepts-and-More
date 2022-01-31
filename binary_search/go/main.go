package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	array := []int{}

	for i := 0; i < 10000000; i++ {
		array = append(array, rand.Intn(100))
	}
	sort.Ints(array)
	beggining := time.Now()
	fmt.Println(binarySearch(array, 16))
	end := time.Now()
	fmt.Println(end.Sub(beggining))

}

// returns the position of a given element in an ordered array
func binarySearch(list []int, elem int) int {

	beggining := 0
	end := len(list) - 1

	for beggining <= end {

		mid := beggining + (end-beggining)/2
		mid_value := list[mid]
		//fmt.Println(beggining, mid, end)
		if mid_value == elem {
			return mid
		} else if elem < mid_value {

			end = mid - 1
		} else {
			beggining = mid + 1
		}

	}

	fmt.Println("Item not in list")
	return 0

}

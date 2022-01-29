package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	// This scope is gone each time the function is called
	a := 0 // two prior
	b := 0 // one prior
	c := 0 // current
	count := 0

	return func() int {
		// This one still exists for some magical reason
		count += 1
		if count == 1 {
			a = 0
			b = 0
			c = a + b
		} else if count == 2 {
			a = 0
			b = 1
			c = a + b
		} else {
			fmt.Printf("a: %d; b: %d; c: %d\n", a, b, c)
			b = c
			c = a + b
			a = b

		}

		return c
	}
}

// 0 1 1 2 3 5 8 13 21 34
func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		//fmt.Printf("%d => %d\n", i, f())
		f()
	}
}

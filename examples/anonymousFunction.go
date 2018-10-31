package main

import (
	"fmt"
)

func main() {

	// Anonymous function
	addFunc := func(x ...int) (terms int, sum int) {
		terms = len(x)
		for _, t := range x {
			sum += t
		}
		return
	}

	terms, sum := addFunc(3, 4, 5, 6, 7)
	fmt.Println("Added ", terms, "terms to get total: ", sum)
}

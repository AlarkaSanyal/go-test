package main

import (
	"fmt"
)

func main() {

	// Return single value
	result := add(1, 3, 5, 7)
	fmt.Println("Sum: ", result)

	// Return multiple values (ODD)
	terms1, sum1, evenOdd1 := addMulti(1, 3, 5, 7, 9)
	fmt.Println("Added ", terms1, "terms(", evenOdd1, ") to get total: ", sum1)

	// Return multiple values (EVEN)
	terms2, sum2, evenOdd2 := addMulti(1, 3, 5, 7, 9, 10)
	fmt.Println("Added ", terms2, "terms(", evenOdd2, ") to get total: ", sum2)

	// Using named variables
	terms3, sum3, evenOdd3 := addNamed(1, 3, 5, 6)
	fmt.Print("Using named -> ")
	fmt.Println("Added ", terms3, "terms(", evenOdd3, ") to get total: ", sum3)
}

func add(x ...int) int {
	result := 0
	for _, term := range x {
		result += term
	}
	return result
}

func addMulti(x ...int) (int, int, string) {
	result := 0
	for _, term := range x {
		result += term
	}
	evenOdd := "ODD"
	if len(x)%2 == 0 {
		evenOdd = "EVEN"
	}
	return len(x), result, evenOdd
}

func addNamed(x ...int) (terms3 int, sum3 int, evenOdd3 string) {
	terms3 = len(x)
	for _, term := range x {
		sum3 += term
	}
	evenOdd3 = "ODD"
	if terms3%2 == 0 {
		evenOdd3 = "EVEN"
	}
	return
}

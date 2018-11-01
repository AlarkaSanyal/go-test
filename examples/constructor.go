package main

import "fmt"

func main() {
	// Call constructor to initialize the map before adding key-values to the map
	foo := newMyStruct()
	foo.myMap[3] = "three"

	fmt.Println("Struct: ", foo)
	fmt.Println("Value of field 1: ", foo.myMap[1])
	fmt.Println("Value of field 3: ", foo.myMap[3])
}

type myStruct struct {
	myMap map[int]string
}

// Constructors generally start with "new" followed by the struct name
// Return the pointer (by adding "*")
func newMyStruct() *myStruct {
	result := myStruct{}
	result.myMap = map[int]string{}
	return &result
}

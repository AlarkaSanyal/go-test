package main

import "fmt"

func main() {

	// Option 1
	fmt.Println("-------Option 1---------")
	st1 := myStruct{} // Creates the field in the local execution stack and not in heap
	st1.field1 = "foo1"
	fmt.Println(st1.field1)

	// Option 2
	fmt.Println("-------Option 2 (Composite literal syntax)---------")
	st2 := myStruct{"foo2"} // Creates the field in the local execution stack and not in heap
	fmt.Println(st2.field1)
	fmt.Println("Composite literal of st2: ", st2)

	// Option 3
	fmt.Println("-------Option 3(reference type using \"new\" function )---------")
	st3 := new(myStruct) // Creates the field in the heap
	st3.field1 = "foo3"
	fmt.Println(st3.field1)
	println("Memory addess of st3: ", st3)
	fmt.Println("Memory addess of st3 (using fmt): ", st3)

	// Option 4
	fmt.Println("-------Option 4(reference type using composite literal syntax)---------")
	st4 := &myStruct{"foo4"} // Creates the field in the heap
	fmt.Println(st4.field1)
	println("Memory addess of st4: ", st4)
	fmt.Println("Memory addess of st4 (using fmt): ", st4)
}

type myStruct struct {
	field1 string
}

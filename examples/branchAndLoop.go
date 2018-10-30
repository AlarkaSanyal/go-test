package main

func main() {

	// If-then-else Statement
	//foo1 := 2

	if foo1 := 2; foo1 == 3 {
		println("Yes")
	} else {
		println("No")
	}

	// Switch Statement
	foo2 := 5

	switch {
	case foo2 == 1:
		println("One")
	case foo2 > 3:
		println("More than 3")
	}

	//Loops

	// For_a
	for i := 0; i < 5; i++ {
		println("For_a: ", i)
	}

	// For_b
	i := 0
	for {
		println("For_b: ", i)
		i++

		if i > 3 {
			break
		}
	}

	// For_slice
	s := []string{"foo", "bar", "buz"}

	for index, value := range s {
		println("For_slice", index, value)
	}

	// For_map
	m := make(map[string]string)
	m["first"] = "foo"
	m["second"] = "second"
	m["third"] = "third"

	for key, value := range m {
		println("For_map: ", key, value)
	}
}

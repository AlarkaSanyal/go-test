package main

import "fmt"

func main() {
	pm := printMessage{"hello"}

	// Without methods
	fmt.Println("Without methods: ", pm.message)

	// Using methods
	pm.printMsg()

}

type printMessage struct {
	message string
}

func (msg *printMessage) printMsg() {
	fmt.Println("Using methods: ", msg.message)
}

package main

import "fmt"

func main() {
	pm := printMessage{"Hello World parent1!"}
	pm.printMsg()

	cpm := childPrintMessage{}
	cpm.message = "Hello World parent2!"
	cpm.extraMessage = "Hello World child!"
	cpm.printChildMsg()
}

type printMessage struct {
	message string
}

func (msg *printMessage) printMsg() {
	fmt.Println("Parent class call: ", msg.message)
}

type childPrintMessage struct {
	printMessage
	extraMessage string
}

func (msg *childPrintMessage) printChildMsg() {
	fmt.Println("Child class call: ", msg.message, " & ", msg.extraMessage)
}

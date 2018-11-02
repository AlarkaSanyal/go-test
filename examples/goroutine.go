package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	// By default, go program runs on only one CPU.
	// Using GOMAXPROCS, we can set the max number of CPUs
	runtime.GOMAXPROCS(8)

	// 'go' (goroutine) keword denotes that the function abcGen() is running Asynchronously (concurrently)
	go abcGen()

	// This is printed before the execution of the goroutine abcGen() function
	// because the current function (in this case main()) is executed first before
	// it either terminates or pauses
	// NB: To understand this, comment out GOMAXPROCS so that only 1 CPU is used
	fmt.Println("This comes first although written after abcGen() function call in the code")

	// Without a sleep, the code exists the main() function before the goroutine can be executed
	time.Sleep(100 * time.Millisecond)
}

func abcGen() {
	fmt.Println("--------------Print A-Z sequentially--------------")
	for l := byte('A'); l <= byte('Z'); l++ {
		fmt.Println(string(l))
	}
	fmt.Println("")

	fmt.Println("--------------Prints A-Z randomly--------------")
	for l := byte('A'); l <= byte('Z'); l++ {
		// This would lead to 26 goroutines and with multiple CPUs working
		// concurrently, the values might not be printed in order.
		go fmt.Println(string(l))
	}
	fmt.Println("")
}

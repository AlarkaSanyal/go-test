package main

func main() {
	// Passing value
	message := "Passing"

	passValue(message)
	println("After: ", message)

	// Passing reference
	passRef(&message)
	println("After: ", message)

	passLots("hello", "world", "lots", "of", "values")
}

// Function for passing value
func passValue(message string) {
	println("Before: ", message)
	// updaitng value to check if main function value changes or not
	message = "new Passing value"
}

// Function for passing reference
func passRef(message *string) {
	println("Before: ", *message)
	// updaitng value to check if main function value changes or not
	*message = "new Passing reference"
}

// Variatic function
func passLots(messages ...string) {
	println("Variatic function output:")
	for _, message := range messages {
		println(message)
	}
}

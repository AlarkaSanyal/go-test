package main

import "fmt"

func main() {

	plantCapacities := []float64{30, 30, 30, 60, 60, 100}
	activePlants := []int{0, 1}
	gridLoad := 75.

	fmt.Println("1) Gen plant report")
	fmt.Println("2) Gen grid report")
	fmt.Println("Choose an option: ")

	var option string

	fmt.Scanln(&option)

	switch {
	case option == "1":
		plantCap(plantCapacities...)

	case option == "2":
		plantDetails(activePlants, plantCapacities, gridLoad)

	default:
		fmt.Println("Unknown option")
	}
}

func plantCap(plantCapacities ...float64) {
	for index, value := range plantCapacities {
		fmt.Printf("Plant %d capacity: %.0f\n", index, value)
	}
}

func plantDetails(activePlants []int, plantCapacities []float64, gridLoad float64) {
	capacity := 0.
	for _, plantID := range activePlants {
		capacity += plantCapacities[plantID]
	}

	fmt.Printf("%-20s%.0f\n", "Capacity: ", capacity)
	fmt.Printf("%-20s%.0f\n", "GridLoad: ", gridLoad)
	fmt.Printf("%-20s%.1f%%\n", "Utilization: ", gridLoad/capacity*100)
}

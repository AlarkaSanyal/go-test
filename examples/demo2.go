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
		for index, value := range plantCapacities {
			fmt.Printf("Plant %d capacity: %.0f\n", index, value)
		}

	case option == "2":
		capacity := 0.
		for _, plantId := range activePlants {
			capacity += plantCapacities[plantId]
		}

		fmt.Printf("%-20s%.0f\n", "Capacity: ", capacity)
		fmt.Printf("%-20s%.0f\n", "GridLoad: ", gridLoad)
		fmt.Printf("%-20s%.1f%%\n", "Utilization: ", gridLoad/capacity*100)

	default:
		fmt.Println("Unknown option")
	}
}
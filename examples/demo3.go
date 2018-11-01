package main

import (
	"fmt"
	"strings"
)

func main() {

	plants := []PowerPlant{
		PowerPlant{hydro, 300, active},
		PowerPlant{wind, 30, active},
		PowerPlant{wind, 25, inactive},
		PowerPlant{wind, 35, active},
		PowerPlant{solar, 45, unavailable},
		PowerPlant{solar, 40, active},
	}

	grid := PowerGrid{300, plants}

	fmt.Println("1) Gen plant report")
	fmt.Println("2) Gen grid report")
	fmt.Println("Choose an option: ")

	var option string

	fmt.Scanln(&option)

	switch {
	case option == "1":
		grid.generatePlantReport()

	case option == "2":
		grid.generateGridReport()

	default:
		fmt.Println("Unknown option")
	}
}

// PlantType represents types of powerplants
type PlantType string

const (
	hydro PlantType = "Hydro"
	wind  PlantType = "Wind"
	solar PlantType = "Solar"
)

// PlantStatus represents status of powerplants
type PlantStatus string

const (
	active      PlantStatus = "Active"
	inactive    PlantStatus = "Inactive"
	unavailable PlantStatus = "Unavailable"
)

// PowerPlant object structure
type PowerPlant struct {
	plantType   PlantType
	capacity    float64
	plantStatus PlantStatus
}

// PowerGrid object structure
type PowerGrid struct {
	load   float64
	plants []PowerPlant
}

func (pg *PowerGrid) generatePlantReport() {
	for index, plant := range pg.plants {
		label := fmt.Sprintf("%s%d", "Plant #", index)
		fmt.Println(label)
		fmt.Println(strings.Repeat("-", len(label)))
		fmt.Printf("%-20s%s\n", "Type: ", plant.plantType)
		fmt.Printf("%-20s%.0f\n", "Capacity: ", plant.capacity)
		fmt.Printf("%-20s%s\n", "Status: ", plant.plantStatus)
		fmt.Println("")
	}
}

func (pg *PowerGrid) generateGridReport() {
	capacity := 0.

	for _, plant := range pg.plants {
		if plant.plantStatus == active {
			capacity += plant.capacity
		}
	}
	label := "Power Grid report"
	fmt.Println(label)
	fmt.Println(strings.Repeat("-", len(label)))
	fmt.Printf("%-20s%.0f\n", "Capacity: ", capacity)
	fmt.Printf("%-20s%.0f\n", "GridLoad: ", pg.load)
	fmt.Printf("%-20s%.1f%%\n", "Utilization: ", pg.load/capacity*100)
	fmt.Println("")
}

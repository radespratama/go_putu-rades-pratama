package main

import "fmt"

type Car struct {
	CarType string
	FuelIn  float64
}

func (c Car) EstimateDistance() float64 {
	var fuelEfficiency float64 = 1.5

	estimatedDistance := c.FuelIn * fuelEfficiency
	return estimatedDistance
}

func main() {
	car := Car{
		CarType: "SUV",
		FuelIn:  10.5,
	}

	estimatedDistance := car.EstimateDistance()
	fmt.Printf("Tipe mobil: %s, Jarak: %.2f\n", car.CarType, estimatedDistance)
}

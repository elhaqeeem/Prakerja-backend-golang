package main

import (
	"fmt"
)

// Struct untuk Car
type Car struct {
	Type    string
	FuelIn  float64
}

// Method untuk menghitung perkiraan jarak yang dapat ditempuh
func (c *Car) EstimateDistance() float64 {
	// Konversi bahan bakar (dalam liter) ke dalam milliliter
	fuelInMilliliters := c.FuelIn * 1000

	// Menghitung perkiraan jarak yang dapat ditempuh
	// dengan asumsi 1.5 L / milliliter
	distance := fuelInMilliliters / 1.5

	return distance
}

func main() {
	// Membuat objek Car
	myCar := Car{
		Type:   "Sedan",
		FuelIn: 30.0, // Misalnya, ada 30 liter bahan bakar dalam tangki
	}

	// Menggunakan method untuk menghitung perkiraan jarak
	estimatedDistance := myCar.EstimateDistance()

	// Menampilkan hasil perkiraan jarak
	fmt.Printf("Perkiraan jarak yang dapat ditempuh dengan %s: %.2f kilometer\n", myCar.Type, estimatedDistance)
}

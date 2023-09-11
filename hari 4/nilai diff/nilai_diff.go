package main

import "fmt"

// Function untuk menemukan nilai maksimum dan minimum
func findMaxMin(numbers *[6]int) (int, int) {
	if numbers == nil {
		return 0, 0
	}

	max := numbers[0]
	min := numbers[0]

	for _, num := range *numbers {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	return max, min
}

func main() {
	var inputNumbers [6]int

	// Meminta pengguna untuk memasukkan 6 angka
	fmt.Println("Masukkan 6 angka:")
	for i := 0; i < 6; i++ {
		fmt.Printf("Angka ke-%d: ", i+1)
		fmt.Scan(&inputNumbers[i])
	}

	// Memanggil function untuk mencari nilai maksimum dan minimum
	max, min := findMaxMin(&inputNumbers)

	// Menampilkan hasil
	fmt.Printf("Nilai maksimum: %d\n", max)
	fmt.Printf("Nilai minimum: %d\n", min)
}

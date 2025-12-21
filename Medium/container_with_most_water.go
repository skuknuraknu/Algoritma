package main

import (
	"fmt"
)

func main() {
	heights := []int{1, 8, 1, 2, 2, 3, 5}
	result1 := maxArea(heights)
	result2 := maxAreaBruteforce(heights)
	fmt.Printf("Hasil Two Pointers: %d\n", result1)
	fmt.Printf("Hasil Brute Force: %d\n", result2)
}

func maxArea(height []int) int {
	kiri, kanan := 0, len(height)-1
	areaTerbesar := 0

	for kiri < kanan {
		jarak := kanan - kiri
		tinggiTerkecil := min(height[kiri], height[kanan])
		areaTerbesar = max(areaTerbesar, jarak*tinggiTerkecil)

		if height[kiri] < height[kanan] {
			kiri++
		} else {
			kanan--
		}
	}
	return areaTerbesar
}
func maxAreaBruteforce(height []int) int {
	maxArea := 0
	ukuranArray := len(height)

	// Loop untuk setiap pasangan indeks
	for kiri := 0; kiri < ukuranArray; kiri++ {
		for kanan := kiri + 1; kanan < ukuranArray; kanan++ {
			// Hitung lebar dan tinggi
			width := kanan - kiri
			tinggiTerkecil := min(height[kiri], height[kanan])
			maxArea = max(maxArea, width*tinggiTerkecil)
		}
	}
	return maxArea
}

package main

func maxArea(height []int) int {
	indexKiri, indexKanan := 0, len(height)-1
	areaTerbesar := 0

	for indexKiri < indexKanan {
		jarak := indexKanan - indexKiri
		tinggiTerkecil := min(height[indexKiri], height[indexKanan])
		areaTerbesar = max(areaTerbesar, jarak*tinggiTerkecil)

		if height[indexKiri] < height[indexKanan] {
			indexKiri++
		} else {
			indexKanan--
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

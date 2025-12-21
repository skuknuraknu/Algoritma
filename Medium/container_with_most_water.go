package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxArea([]int{6, 2, 7, 9, 5, 1})) // 49
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

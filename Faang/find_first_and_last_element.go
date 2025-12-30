package main

import "fmt"

/*
Find First and Last Position of Element in Sorted Array

Diberikan array yang sudah terurut dan nilai target, temukan indeks pertama dan terakhir
dari elemen target tersebut. Jika tidak ditemukan, kembalikan [-1, -1].

Contoh: nums = [5,7,7,8,8,10], target = 8 → hasil = [3,4]

Solusi: Menggunakan binary search dua kali untuk mencapai kompleksitas waktu O(log n),
lebih efisien dibandingkan linear search O(n).
*/
func findFirstAndLastElement(arr []int, target int) (int, int) {
	fmt.Printf("\n=== Mencari target %d dalam array %v ===\n", target, arr)
	kiri := cariTerendah(arr, target)
	if kiri == len(arr) || arr[kiri] != target {
		return -1, -1 // Target tidak ditemukan
	}
	kanan := cariTertinggi(arr, target) - 1 // Dikurangi 1 karena cariTertinggi mengembalikan indeks setelah elemen terakhir
	return kiri, kanan
}

func cariTerendah(nums []int, target int) int {
	kiri, kanan := 0, len(nums)

	for kiri < kanan {
		tengah := kiri + (kanan-kiri)/2 // Menghindari overflow

		if nums[tengah] >= target {
			// Jika nilai tengah >= target, kemungkinan masih ada elemen yang sama di sebelah kiri
			// Pindahkan pointer kanan ke posisi tengah untuk melanjutkan pencarian ke kiri
			kanan = tengah
		} else {
			// Jika nilai tengah < target, target pasti berada di sebelah kanan
			// Pindahkan pointer kiri ke tengah+1
			kiri = tengah + 1
		}
	}

	return kiri // Mengembalikan indeks pertama yang >= target
}

func cariTertinggi(nums []int, target int) int {
	kiri, kanan := 0, len(nums)

	for kiri < kanan {
		tengah := kiri + (kanan-kiri)/2 // Menghindari overflow

		if nums[tengah] > target {
			// Jika nilai tengah > target, batas akhir berada di sebelah kiri
			kanan = tengah
		} else {
			// Jika nilai tengah <= target, kemungkinan masih ada elemen yang sama di sebelah kanan
			// Pindahkan pointer kiri ke tengah+1 untuk melanjutkan pencarian ke kanan
			kiri = tengah + 1
		}
	}

	return kiri // Mengembalikan indeks setelah elemen target terakhir
}

package main

/*
Valid Mountain Array
Given an array of integers arr, return true if and only if it is a valid mountain array.

An array is a mountain array if and only if:
- arr.length ≥ 3
- There exists some i with 0 < i < arr.length - 1 such that:
	- arr[0] < arr[1] < ... < arr[i - 1] < arr[i]
	- arr[i] > arr[i + 1] > ... > arr[arr.length - 1]
*/
func validMountainArray(arr []int) bool {
	n := len(arr)
	if n < 3 {
		return false
	}
	index := 0

	// Mencoba naik puncak
	for index+1 < n && arr[index] < arr[index+1] {
		index++
		println("Naik ke index :", index)
	}
	// Puncak tidak boleh di awal/akhir
	if index == 0 || index == n-1 {
		println("Puncak tidak valid di index :", index)
		return false
	}
	println("Puncak di index :", index)
	for index+1 < n && arr[index] > arr[index+1] {
		index++
		println("Turun ke index :", index)
	}
	println("Akhir di index :", index)
	return index == n-1
}

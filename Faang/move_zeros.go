package main

func moveZeros(nums []int) []int {
	writePosition := 0
	for _, val := range nums {
		if val != 0 {
			nums[writePosition] = val
			writePosition++
		}
	}
	for i := writePosition; i < len(nums); i++ {
		nums[i] = 0
	}
	return nums
}

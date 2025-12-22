package main

import (
	"sort"
)

/*
You are given an array people where people[i] is the weight of the ith person, and an infinite number of boats where each boat can carry a maximum weight of limit.

Each boat carries at most two people at the same time, provided the sum of the weight of those people is at most limit.

Return the minimum number of boats to carry every given person.
*/
func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	indexKiri, indexKanan := 0, len(people)-1
	jumlahKapal := 0
	for indexKiri <= indexKanan {
		if people[indexKiri]+people[indexKanan] <= limit {
			indexKiri++
		}
		indexKanan--
		jumlahKapal++
	}
	return jumlahKapal
}

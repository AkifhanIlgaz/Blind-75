package main

import "math"

func main() {

}

func longestConsecutive(nums []int) int {
	set := map[int]bool{}

	var longest float64
	for _, num := range nums {
		set[num] = true
	}

	for _, num := range nums {
		if !set[num - 1] {
			current := 0

			for set[num+current] {
				current++
			}

			longest = math.Max(float64(current), float64(longest))
		}
	}

	return int(longest)
}
package main

import "sort"

func main() {
}

func topKFrequent(nums []int, k int) []int {

	rec := make(map[int]int, len(nums))

	for _, num := range nums {
		rec[num]++
	}	

	counts := make([]int, 0, len(rec))

	for _, count := range rec {
		counts = append(counts, count)
	}

	sort.Ints(counts)

	result := make([]int, 0, k)
	min := counts[len(counts) - k]

	for num, count := range rec {
		if count >= min {
			result = append(result, num)
		}
	}

	return result

}
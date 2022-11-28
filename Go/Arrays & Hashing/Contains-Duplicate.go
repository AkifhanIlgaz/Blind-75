package main

func main() {}

func containsDuplicate(nums []int) bool {
	set := map[int]int{}

	for _, num := range nums {
		if set[num] != 0 {
			return true
		} else {
			set[num] = 1
		}
	}

	return false
}
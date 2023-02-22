package main

func main() {

}

func containsDuplicate(nums []int) bool {
	// empty struct has zero size
	numSet := make(map[int]struct{})

	for _, num := range nums {
		if _, found := numSet[num]; found {
			return true
		} else {
			numSet[num] = struct{}{}
		}
	}

	return false
}

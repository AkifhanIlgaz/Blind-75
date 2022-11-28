package main

func main() {}

func twoSum(nums []int, target int) []int {
	prevMap := map[int]int{} // val => index

	for index, val := range nums {
		diff := target - val
		if value, ok := prevMap[diff]; ok {
			return []int{value, index}
		}

		prevMap[val] = index
	}

	return []int{}
}
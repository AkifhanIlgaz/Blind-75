package main

func main() {

}

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums), len(nums))
	for index := range result {
		result[index] = 1
	}

	prefix := 1

	for i := 0; i < len(nums); i++ {
		result[i] = prefix
		prefix *= nums[i]
	}

	postfix := 1

	for i := len(nums) -1 ; i >= 0; i-- {
		result[i] *= postfix
		postfix *= nums[i]
	}

	return result

}
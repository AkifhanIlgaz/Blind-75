package main

func main() {

}

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)

	for currentIndex, currentVal := range nums {
		difference := target - currentVal
		if requiredIndex, isPresent := numsMap[difference]; isPresent {
			return []int{currentIndex, requiredIndex}
		}
		numsMap[currentVal] = currentIndex
	}

	return []int{}
}

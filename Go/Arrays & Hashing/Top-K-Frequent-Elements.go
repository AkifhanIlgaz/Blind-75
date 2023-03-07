package main

func main() {

}

func topKFrequent(nums []int, k int) []int {
	frequencyMap := make(map[int]int, len(nums))
	frequencySlice := make([][]int, len(nums)+1)

	for _, num := range nums {
		frequencyMap[num]++
	}

	for num, count := range frequencyMap {
		frequencySlice[count] = append(frequencySlice[count], num)
	}

	
	result := []int{}

	for i := len(frequencySlice) - 1; i > 0; i-- {
		result = append(result, frequencySlice[i]...)
		if len(result) == k {
			break
		}
	}

	return result
}

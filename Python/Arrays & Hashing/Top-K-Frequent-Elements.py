class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        frequencyMap = {}
        frequencyList = [[] for i in range(len(nums) + 1)]

        for num in nums:
            frequencyMap[num] = frequencyMap.get(num, 0) + 1

        for num, count in frequencyMap.items():
            frequencyList[count].append(num)

        result = []

        for i in range(len(frequencyList) - 1, 0, -1):
            for n in frequencyList[i]:
                result.append(n)
                if len(result) == k:
                    return result

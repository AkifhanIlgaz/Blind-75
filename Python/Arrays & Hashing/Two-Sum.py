class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        numsMap = {}

        for currentIndex, currentVal in enumerate(nums):
            difference = target - currentVal
            if difference in numsMap:
                return [currentIndex, numsMap[difference]]

            numsMap[currentVal] = currentIndex

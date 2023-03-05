/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function (nums, target) {
  const numsMap = new Map();

  for (currentIndex in nums) {
    const currentVal = nums[currentIndex];
    const difference = target - currentVal;
    if (numsMap.has(difference)) {
      return [currentIndex, numsMap.get(difference)];
    }

    numsMap.set(currentVal, currentIndex);
  }
};

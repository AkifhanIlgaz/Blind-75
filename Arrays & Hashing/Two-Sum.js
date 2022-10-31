/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
 var twoSum = function(nums, target) {
    let targetMap = new Map();
  
    for (let index in nums) {
      if (targetMap.has(target - nums[index])) {
        return [index, targetMap.get(target - nums[index])];
      }
      targetMap.set(nums[index], index);
    }
  }
  
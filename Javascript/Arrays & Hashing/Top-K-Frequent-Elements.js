/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number[]}
 */
var topKFrequent = function (nums, k) {
  const frequencyMap = new Map();
  const frequencyArray = new Array(nums.length + 1).fill([]);

  nums.forEach((num) =>
    frequencyMap.has(num)
      ? frequencyMap.set(num, frequencyMap.get(num) + 1)
      : frequencyMap.set(num, 1)
  );

  for (let [num, count] of frequencyMap.entries()) {
    frequencyArray[count] = [...frequencyArray[count], num];
  }

  const result = [];

  for (let i = frequencyArray.length - 1; i > 0; i--) {
    result.push(...frequencyArray[i]);
    if (result.length == k) {
      break;
    }
  }

  return result;
};

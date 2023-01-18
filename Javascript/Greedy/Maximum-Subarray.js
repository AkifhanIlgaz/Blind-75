/**
 *
 * @param {Number[]} nums
 * @return {Number}
 */
const maxSubArray = (nums) => {
  let currentSum = 0;
  let maximumSum = nums[0];

  for (let i = 0; i < nums.length; i++) {
    if (currentSum < 0) {
      currentSum = 0;
    }

    currentSum += nums[i];

    maximumSum = Math.max(maximumSum, currentSum);
  }

  return maximumSum;
};

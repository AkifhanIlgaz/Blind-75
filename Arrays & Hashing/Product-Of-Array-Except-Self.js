/**
 * @param {number[]} nums
 * @return {number[]}
 */
const productExceptSelf = (nums) => {
  let prefix = new Array(nums.length),
    postfix = new Array(nums.length);

  prefix[0] = 1;
  for (let i = 1; i < nums.length; i++) {
    prefix[i] = prefix[i - 1] * nums[i - 1];
  }

  postfix[postfix.length - 1] = 1;
  for (let i = postfix.length - 2; i >= 0; i--) {
    postfix[i] = postfix[i + 1] * nums[i + 1];
  }

  return prefix.map((element, index) => element * postfix[index]);
};

/**
 * @param {number[]} nums
 * @return {number[]}
 * Time O(N) | Space O(1)
 */

const productExceptSelf1 = (nums) => {
  let result = new Array(nums.length).fill(1);

  let prefix = 1;

  for (let i = 0; i < nums.length; i++) {
    result[i] = prefix;
    prefix *= nums[i];
  }

  let postfix = 1;

  for (let i = result.length - 1; i >= 0; i--) {
    result[i] *= postfix;
    postfix *= nums[i];
  }

  return result;
};

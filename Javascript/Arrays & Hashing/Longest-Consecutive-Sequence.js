/**
 * @param {number[]} nums
 * @return {number}
 */
const longestConsecutive = (nums) => {
  let set = new Set(nums);

  let longestConsecutive = 0;

  for (let number of nums) {
    if (!set.has(number - 1)) {
      let current = 0;
      while (set.has(number + current)) {
        current++;
      }

      longestConsecutive = Math.max(current, longestConsecutive);
    }
  }

  return longestConsecutive;
};

/**
 *
 * @param {string[]} strs
 * @return {string[][]}
 */
const groupAnagrams = (strs) => {
  let sortedStrings = new Map();

  strs.forEach((str) => {
    const sorted = str.split("").sort().join("");

    if (sortedStrings.has(sorted)) {
      const prev = sortedStrings.get(sorted);
      prev.push(str);
      sortedStrings.set(sorted, prev);
    } else {
      sortedStrings.set(sorted, [str]);
    }
  });

  return Array.from(sortedStrings.values());
};

groupAnagrams(["eat", "tea", "tan", "ate", "nat", "bat"]);

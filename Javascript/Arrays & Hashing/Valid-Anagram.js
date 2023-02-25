/**
 *
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */

const isAnagram = (s, t) => {
  if (s.length != t.length) return false;

  const frequency = new Array(26).fill(0);

  for (let i = 0; i < s.length; i++) {
    const [sChar, tChar] = [s[i].charCodeAt(0), t[i].charCodeAt(0)];
    const a = "a".charCodeAt(0);

    frequency[sChar - a]++;
    frequency[tChar - a]--;
  }

  return frequency.every((freq) => freq == 0);
};

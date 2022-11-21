/**
 *
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
const isAnagram = (s, t) => {
  if (s.length !== t.length) return false;

  let mapS = {};
  let mapT = {};

  for (let i = 0; i < s.length; i++) {
    let charS = s[i];
    let charT = t[i];

    mapS[charS] ? mapS[charS]++ : (mapS[charS] = 1);
    mapT[charT] ? mapT[charT]++ : (mapT[charT] = 1);
  }

  for (let letter in mapS) {
    if (mapS[letter] !== mapT[letter]) return false;
  }

  return true;
};

console.log(isAnagram("aa", "bb"));

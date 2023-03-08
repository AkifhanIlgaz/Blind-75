/**
 * @param {string} s
 * @return {boolean}
 */
const isPalindrome = (s) => {
  let left = 0;
  let right = s.length - 1;

  while (left < right) {
    let [leftChar, rightChar] = [s[left].toLowerCase(), s[right].toLowerCase()];
    if (isNonAlphaNumeric(leftChar)) {
      left++;
      continue;
    }
    if (isNonAlphaNumeric(rightChar)) {
      right--;
      continue;
    }

    if (leftChar != rightChar) {
      return false;
    }

    left++;
    right--;
  }

  return true;
};

const isNonAlphaNumeric = (char) => {
  const isNonAlpha = char < "a" || "z" < char; // a(97) - z(122)
  const isNonNumeric = char < "0" || "9" < char; // 0(48) - 9(57)

  return isNonAlpha && isNonNumeric;
};

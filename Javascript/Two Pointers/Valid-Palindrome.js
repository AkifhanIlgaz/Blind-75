/**
 * @param {string} s
 * @return {boolean}
 */
const isPalindrome = (s) => {
  if (s.length == 0) return true;

  s = s.toLowerCase();

  return isValid(s);
};

/**
 *
 * @param {string} s
 * @return {boolean}
 */
const isValid = (s) => {
  let [leftPointer, rightPointer] = [0, s.length - 1];

  while (leftPointer < rightPointer) {
    while (leftPointer < rightPointer && isNonAlphaNumeric(s[leftPointer])) {
      leftPointer++;
    }
    while (leftPointer < rightPointer && isNonAlphaNumeric(s[rightPointer])) {
      rightPointer--;
    }

    if (s[leftPointer] !== s[rightPointer]) return false;

    leftPointer++;
    rightPointer--;
  }

  return true;
};

const isNonAlphaNumeric = (char) => {
  const isNonAlpha = char < "a" || "z" < char; // a(97) - z(122)
  const isNonNumeric = char < "0" || "9" < char; // 0(48) - 9(57)

  return isNonAlpha && isNonNumeric;
};

/**
 *
 * @param {number} n
 * @returns {string[]}
 */
const generateParenthesis = (n) => {
  const stack = [];
  const res = [];

  function backtrack(openN, closedN) {
    if (openN == closedN && closedN == n) {
      res.push(stack.join(""));
      return;
    }

    if (openN < n) {
      stack.push("(");
      backtrack(openN + 1, closedN);
      stack.pop();
    }
    if (closedN < openN) {
      stack.push(")");
      backtrack(openN, closedN + 1);
      stack.pop();
    }
  }

  backtrack(0, 0);
  return res;
};

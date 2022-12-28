/**
 *
 * @param {number[]} temperatures
 * @returns {number[]}
 */
const dailyTemperatures = (temperatures) => {
  const res = new Array(temperatures.length).fill(0);
  const stack = [];

  for (i in temperatures) {
    const t = temperatures[i];
    while (stack.length > 0 && t > stack[stack.length - 1][0]) {
      let [stackT, stackInd] = stack.pop();
      res[stackInd] = i - stackInd;
    }
    stack.push([t, i]);
  }
  return res;
};

let INT_BITS = 32;

function rightRotate(n, d) {
  /* In n>>d, first d bits are 0.
    To put last 3 bits of at
    first, do bitwise or of n>>d
    with n <<(INT_BITS - d) */
  return (n >> d) | (n << (INT_BITS - d));
}

console.log(rightRotate(6, 26));

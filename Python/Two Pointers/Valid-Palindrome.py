class Solution:
    def isPalindrome(self, s: str) -> bool:
        left, right = 0, len(s) - 1

        while left < right:
            left_char, right_char = s[left].lower(), s[right].lower()
            if not self.alphanum(left_char):
                left += 1
                continue
            if not self.alphanum(right_char):
                right -= 1
                continue

            if left_char != right_char:
                return False

            left += 1
            right -= 1

        return True

    def alphanum(self, c):
        return (
            ord("A") <= ord(c) <= ord("Z")
            or ord("a") <= ord(c) <= ord("z")
            or ord("0") <= ord(c) <= ord("9")
        )

class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False

        frequency = [0] * 26

        for (s_char, t_char) in zip(s, t):
            frequency[ord(s_char) - ord("a")] += 1
            frequency[ord(t_char) - ord("a")] -= 1

        return len(set(frequency)) == 1

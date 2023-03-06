class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        anagrams = collections.defaultdict(list)

        for str in strs:
            counter = [0] * 26
            for char in str:
                counter[ord(char) - ord("a")] += 1
            anagrams[tuple(counter)].append(str)

        return anagrams.values()

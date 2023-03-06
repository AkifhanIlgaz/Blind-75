package main

func main() {

}

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[[26]int][]string)

	for _, str := range strs {
		counter := [26]int{}

		for _, char := range str {
			counter[char-'a']++
		}

		anagrams[counter] = append(anagrams[counter], str)
	}

	values := [][]string{}

	for _, sameAnagrams := range anagrams {
		values = append(values, sameAnagrams)
	}

	return values
}

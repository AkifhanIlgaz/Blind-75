package main

import (
	"sort"
	"strings"
)

func main() {

}



func groupAnagrams(strs []string) [][]string {
	// anagrams := map[string][]string{}

	// for _, s := range strs {
	// 	sorted := SortString(s)
		
	// 	if _,ok := anagrams[sorted]; ok {
	// 		anagrams[sorted] = append(anagrams[sorted], s)
	// 	} else {
	// 		anagrams[sorted] = []string{s}
	// 	}
	// }

	// keys := make([][]string, 0)
	// for k := range anagrams {
	// 	keys = append(keys, anagrams[k])
	// }
	
	// return keys

	anagrams := map[[26]int][]string{}

	for _, s := range strs {
		counter := [26]int{}

		for _, char := range s {
			counter[char - 'a']++
		}

		anagrams[counter] = append(anagrams[counter], s)
	}

	values := [][]string{}
	for _,k := range anagrams {
		values = append(values, k)
	}
	
	return values
}

func SortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
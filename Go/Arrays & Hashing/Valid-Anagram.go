package main

import "fmt"

func main() {
	x := "hello"

	for _, char := range x {
		fmt.Println(char)
	}
	isAnagram("anagram",
	"nagaram")
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hashmap := map[int]int{}
	addFrequency(s, hashmap)
	subtractFrequency(t, hashmap)
	return checkFrequency(hashmap)
}

func addFrequency(str string, hashmap map[int]int) {
	for _,char := range str {
		hashmap[int(char)]++
	}
}

func subtractFrequency(str string, hashmap map[int]int) {
	for _,char := range str {
		if hashmap[int(char)] == 0 {
			continue
		}

		hashmap[int(char)]--
	}
}

func checkFrequency(hashmap map[int]int) bool {
	for _, value := range hashmap {
		if value != 0 {
			return false
		}
	}

	return true
}
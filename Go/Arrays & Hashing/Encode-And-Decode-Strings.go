package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(encode([]string{"neet", "co#de", "is", "awesome"}))

	fmt.Println(decode("4#neet5#co#de2#is7#awesome"))
}

func encode(strs []string) string {
	var stringBuilder strings.Builder

	for _, str := range strs {
		decodedWord := fmt.Sprint(len(str), "#", str)
		stringBuilder.WriteString(decodedWord)
	}

	return stringBuilder.String()
}

func decode(str string) []string {
	decodedWords := []string{}
	i := 0

	for i < len(str) {
		j := i
		for str[j] != '#' {
			j++
		}
		len, _ := strconv.Atoi(str[i:j])
		word := str[j+1 : j+1+len]
		decodedWords = append(decodedWords, word)
		i = j + 1 + len
	}

	return decodedWords
}

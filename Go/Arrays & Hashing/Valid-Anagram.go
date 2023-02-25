package main

func main() {

}

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	frequency := [26]int{}

	for i, sChar := range s {
		tChar := t[i]
		frequency[sChar-'a']++
		frequency[tChar-'a']--
	}

	for _, val := range frequency {
		if val != 0 {
			return false
		}
	}

	return true
}

package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	quantityOfLetters := map[rune]int{}

	for _, char := range magazine {
		quantityOfLetters[char]++
	}

	for _, char := range ransomNote {
		quantityOfLetters[char]--
		if quantityOfLetters[char] < 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(canConstruct("abc", "abc"))
}

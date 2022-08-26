package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	stringX := strconv.Itoa(x)
	runeSliceX := []rune(stringX)
	lenSliceX := len(runeSliceX)
	for i, char := range runeSliceX {
		if char != runeSliceX[lenSliceX-i-1] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isPalindrome(-121))
}

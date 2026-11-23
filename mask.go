package main

import (
	"fmt"
	"slices"
)

func makeMask(word string) []rune {
	runes := []rune(word)
	mask := make([]rune, len(runes))
	for i := range mask {
		mask[i] = '*'
	}
	return mask
}

func assertMask(word string, mask []rune) bool {
	wordMask := []rune(word)
	return slices.Equal(wordMask, mask)
}

func printMask(mask []rune) {
	fmt.Println(" " + string(mask))
}

func openThisChar(char rune, word string, mask []rune) (bool, []rune) {
	wordMask := []rune(word)
	isCharInWord := slices.Contains(wordMask, char)
	if !isCharInWord {
		return false, mask
	}
	for i := range mask {
		if char == wordMask[i] {
			mask[i] = wordMask[i]
		}
	}
	return true, mask
}

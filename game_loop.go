package main

func isGameFinished(word string, mask []rune, mistakes int) bool {
	if mistakes <= 0 {
		return true
	}
	isMaskEqual := assertMask(word, mask)
	if isMaskEqual {
		return true
	}
	return false
}

package main

import (
	"fmt"
	"log"
	"os"
)

const maxMistakes int = 5

func main() {
	dictionary, err := readWordsFromFile("words.txt")
	if err != nil {
		log.Fatalf("не удалось прочитать файл: %v", err)
	}

	startGame := printStartMenu()
	if !startGame {
		os.Exit(0)
	}

	mistakes := maxMistakes
	word := getRandomWord(dictionary)
	mask := makeMask(word)
	fmt.Print(word + " ")
	fmt.Print(mistakes)
	printMask(mask)

	char := getCharFromUser()
	fmt.Println(string(char))
}

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	dictionary, err := readWordsFromFile("words.txt")
	if err != nil {
		log.Fatalf("не удалось прочитать файл: %v", err)
	}
	word := getRandomWord(dictionary)
	fmt.Println(word)

	startGame := printStartMenu()
	if !startGame {
		os.Exit(0)
	}

}

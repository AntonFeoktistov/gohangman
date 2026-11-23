package main

import (
	"fmt"
	"log"
)

func main() {
	dictionary, err := readWordsFromFile("words.txt")
	if err != nil {
		log.Fatalf("не удалось прочитать файл: %v", err)
	}
	word := getRandomWord(dictionary)
	fmt.Println(word)
}
